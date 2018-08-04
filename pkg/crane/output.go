package crane

// OutputConfig sontains configuration information for an output
type OutputConfig struct {
	SpinnerIndex    int    `yaml:"spinner_index"`
	SpinnerInterval int    `yaml:"spinner_interval"`
	SpinnerColor    string `yaml:"spinner_color"`
}

// GetOutput returns the configuration information for an output
func (config *Config) GetOutput(name string) *OutputConfig {
	if config.Outputs == nil {
		config.Outputs = make(map[string]*OutputConfig)
	}

	output := config.Outputs[name]
	if output == nil {
		output = &OutputConfig{}
		config.Outputs[name] = output
	}
	return output
}
