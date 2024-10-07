package config

type YamlConfig struct {
	Idle []IdleConfigItem
}

type IdleConfigItem struct {
	Timeout uint32
	Action  IdleAction
}

type IdleAction struct {
	Type     string   `yaml:"type"`
	OnIdle   []string `yaml:"idle,omitempty"`
	OnResume []string `yaml:"resume,omitempty"`
}
