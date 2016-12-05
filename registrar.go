package registrar

import (
	"bytes"
	"encoding/json"
	"net/http"
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
	host   string
	config map[string]string
}

// Init - initialise a new service registrar instance
func Init(host string, config map[string]string) *Registrar {
	return &Registrar{host, config}
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
