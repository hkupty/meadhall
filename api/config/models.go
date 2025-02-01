package config

type YamlConfig struct {
	Idle []IdleConfigItem
}

type IdleConfigItem struct {
	Timeout uint32
	Action  IdleAction
}
