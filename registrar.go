package registrar

import (
	"log"

	"github.com/ewanvalentine/stack-registrar/providers"
	"github.com/ewanvalentine/stack-registrar/services"
)

const (
	defaultHost = "http://kong:8001/apis"
)

type Registry interface {
	Register(service services.Service) error
}

type Registrar struct {
	host     string
	provider providers.Provider
}

type ConfigOptions struct {
	host     string
	provider providers.Provider
}

type ConfigOption func(*ConfigOptions) error

func SetHost(host string) ConfigOption {
	return func(opt *ConfigOptions) error {
		opt.host = host
		return nil
	}
}

func SetProvider(provider providers.Provider) ConfigOption {
	return func(opt *ConfigOptions) error {
		opt.provider = provider
		return nil
	}
}

// Init - initialise a new service registrar instance
func Init(options ...ConfigOption) *Registrar {

	opt := &ConfigOptions{}

	for _, op := range options {
		err := op(opt)
		if err != nil {
			log.Fatalf("Error rendering configuration: %v", err)
		}
	}

	host := defaultHost

	if opt.host != "" {
		host = opt.host
	}

	provider := providers.Kong(host)

	if opt.provider != nil {
		provider = opt.provider
	}

	// @todo - add environment variable check here

	return &Registrar{host, provider}
}

// Register - register service
func (registry *Registrar) Register(service services.Service) error {
	return registry.provider.Register(service)
}
