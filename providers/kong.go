package providers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ewanvalentine/stack-registrar/services"
)

type KongService struct {
	Name        string `json:"name"`
	UpstreamUrl string `json:"upstream_url"`
	Host        string `json:"host"`
}

type KongProvider struct {
	Host string
}

func Kong(host string) Provider {
	return &KongProvider{host}
}

// Register - Register a new service.
func (provider *KongProvider) Register(service services.Service) error {
	kongService := KongService{
		service.Name,
		service.Upstream,
		service.Host,
	}
	return provider.makePostRequest(kongService)
}

func (provider *KongProvider) Resolve(id string) (*services.Service, error) {
	kongService, err := provider.makeGetRequest(id)

	if err != nil {
		return nil, err
	}

	service := &services.Service{
		Name:     kongService.Name,
		Host:     kongService.Host,
		Upstream: kongService.UpstreamUrl,
	}

	return service, nil
}

// makeRequest - Make a request to the service registry
func (provider *KongProvider) makePostRequest(service KongService) error {
	data, _ := json.Marshal(service)
	_, err := http.Post(provider.Host, "application/json", bytes.NewBuffer(data))
	return err
}

// makeGetRequest - Make a get request to Kong
func (provider *KongProvider) makeGetRequest(id string) (*KongService, error) {
	var kongService *KongService
	response, err := http.Get(provider.Host)
	if err != nil {
		return kongService, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&kongService)

	return kongService, err
}
