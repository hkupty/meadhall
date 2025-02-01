package config

import (
	idle "github.com/hkupty/meadhall/pkg/barrels/idle/models"
)

type YamlConfig struct {
	Idle []IdleConfigItem
}

type IdleConfigItem struct {
	Timeout uint32
	Action  idle.IdleAction
}
