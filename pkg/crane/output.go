package crane

// OutputConfig sontains configuration information for an output
type OutputConfig struct {
	SpinnerIndex    int    `json:"spinner_index,omitempty" yaml:"spinner_index,omitempty"`
	SpinnerInterval int    `json:"spinner_interval,omitempty" yaml:"spinner_interval,omitempty"`
	SpinnerColor    string `json:"spinner_color,omitempty" yaml:"spinner_color,omitempty"`
}

/*
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
*/
