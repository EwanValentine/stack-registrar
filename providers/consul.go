package providers

import "github.com/ewanvalentine/stack-registrar/services"

type ConsulService struct {
}

type ConsulProvider struct {
	Host string
}

func Consul(host string) Provider {
	return &ConsulProvider{host}
}

func (provider *ConsulProvider) Register(service services.Service) error {
	return nil
}

func (provider *ConsulProvider) Resolve(id string) (*services.Service, error) {
	return nil, nil
}
