package crane

// ServiceConfig contains configuration information for a service
type ServiceConfig struct {
	Login string `json:"login,omitempty" yaml:"login,omitempty"`
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}

/*
// GetService returns the configuration information for a service
func (config *Config) GetService(name string) *ServiceConfig {
	if config.Services == nil {
		config.Services = make(map[string]*ServiceConfig)
	}

	service := config.Services[name]
	if service == nil {
		service = &ServiceConfig{}
		config.Services[name] = service
	}
	return service
}
*/