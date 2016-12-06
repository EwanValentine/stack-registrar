package registrar

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const (
	defaultHost = "http://localhost:8001"
)

type Service struct {
	Name     string
	Upstream string
	Host     string
}

type Registry interface {
	Register(service Service) error
}

type Registrar struct {
	host string
}

type ConfigOption struct {
	host string
}

type ConfigOptions func(*ConfigOption) error

func SetHost(host string) ConfigOptions {
	return func(opt *ConfigOption) error {
		opt.host = host
		return nil
	}
}

// Init - initialise a new service registrar instance
func Init(options ...ConfigOptions) *Registrar {

	opt := &ConfigOption{}

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

	return &Registrar{host}
}

// Register - Register a new service.
func (registrar *Registrar) Register(service Service) error {
	return registrar.makeRequest(service)
}

// makeRequest - Make a request to the service registry
func (registrar *Registrar) makeRequest(service Service) error {
	data, _ := json.Marshal(service)
	_, err := http.Post(registrar.host, "application/json", bytes.NewBuffer(data))
	return err
}
