package config

type IdleAction struct {
	Type     string   `yaml:"type"`
	OnIdle   []string `yaml:"idle,omitempty"`
	OnResume []string `yaml:"resume,omitempty"`
}
