package services

import (
	"context"
	"crypto/ecdsa"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/eapache/go-resiliency/breaker"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/oceano-dev/microservices-go-common/config"
	"github.com/oceano-dev/microservices-go-common/models"
)

type SecurityKeysService interface {
	GetAllPublicKeys() ([]*models.PublicKey, error)
}

type securityKeysService struct {
	config *config.Config
}

func NewSecurityKeysService(
	config *config.Config,
) *securityKeysService {
	return &securityKeysService{
		config: config,
	}
}

func (s *securityKeysService) GetAllPublicKeys() ([]*models.PublicKey, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := s.requestJWKS(ctx)
	if err != nil {
		return nil, err
	}

	modelsPublicKeys, err := s.getPublicKeysFromDataJWKS(data)
	if err != nil {
		return nil, err
	}

	return modelsPublicKeys, nil
}

func (s *securityKeysService) requestJWKS(ctx context.Context) ([]byte, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	var err error
	request, err := http.NewRequestWithContext(ctx, "GET", s.config.SecurityKeys.EndPointGetPublicKeys, nil)
	if err != nil {
		log.Println("request error:", err)
		return nil, err
	}

	var response *http.Response
	r := retrier.New(retrier.ConstantBackoff(6, 100*time.Millisecond), nil)
	err = r.Run(func() error {
		b := breaker.New(6, 1, 5*time.Second)
		for {
			result := b.Run(func() error {
				response, err = client.Do(request)
				if err != nil {
					return err
				}

				return nil
			})

			switch result {
			case nil:
				return nil
			case breaker.ErrBreakerOpen:
				return err
			default:
				return err
			}
		}
	})

	if err != nil {
		log.Println("response error:", err)
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("data parse:", err)
		return nil, err
	}

	return data, nil
}

func (s *securityKeysService) getPublicKeysFromDataJWKS(data []byte) ([]*models.PublicKey, error) {
	var modelsPublicKeys []*models.PublicKey

	publicKeyParams := make([]models.PublicKeysParams, 0)
	json.Unmarshal([]byte(data), &publicKeyParams)

	for _, model := range publicKeyParams {
		keyByte, err := json.Marshal(model.Params)
		if err != nil {
			log.Printf("failed to json parse params: %s", err)
			return nil, err
		}

		set, err := jwk.Parse(keyByte)
		if err != nil {
			log.Printf("failed to set public key: %s", err)
			return nil, err
		}

		for it := set.Iterate(context.Background()); it.Next(context.Background()); {
			pair := it.Pair()
			key := pair.Value.(jwk.Key)

			var rawkey interface{}
			if err := key.Raw(&rawkey); err != nil {
				log.Printf("failed to create public key: %s", err)
				return nil, err
			}

			publicKey, ok := rawkey.(*ecdsa.PublicKey)
			if !ok {
				log.Printf("expected ecdsa key, got %T", rawkey)
				return nil, err
			}

			modelPublicKey := &models.PublicKey{}
			modelPublicKey.Key = publicKey
			modelPublicKey.Kid = model.Kid
			modelPublicKey.ExpiresAt = model.ExpiresAt
			modelsPublicKeys = append(modelsPublicKeys, modelPublicKey)
		}
	}

	return modelsPublicKeys, nil
}
