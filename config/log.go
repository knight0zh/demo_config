package config

type LogCfg struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxBackups uint   `json:"max_backups"`
}
