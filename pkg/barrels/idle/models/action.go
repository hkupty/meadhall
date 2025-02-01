package models

// github.com/hkupty/meadhall/pkg/barrels/idle/models

type IdleAction struct {
	Type     string   `yaml:"type"`
	OnIdle   []string `yaml:"idle,omitempty"`
	OnResume []string `yaml:"resume,omitempty"`
}
