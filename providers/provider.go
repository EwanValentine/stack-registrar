package providers

import "github.com/ewanvalentine/stack-registrar/services"

type Provider interface {
	Register(service services.Service) error
	Resolve(id string) (*services.Service, error)
}
