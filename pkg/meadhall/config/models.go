package config

type YamlConfig struct {
	Idle []IdleConfigItem
}

type IdleConfigItem struct {
	Timeout uint32
	Action  IdleAction
}

type IdleAction struct {
	Type       string   `yaml:"type"`
	Args       []string `yaml:"args,omitempty"`
	ResumeArgs []string `yaml:"args,omitempty"`
}
