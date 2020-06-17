package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigBuild struct {
	Viper *viper.Viper
	Mysql MysqlClusterCfg `mapstructure:"mysql"`
	Redis RedisClusterCfg `mapstructure:"redis"`
	Log   *LogCfg         `mapstructure:"log"`
}

func NewConfig(path string) *ConfigBuild {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var uluCfg ConfigBuild
	if err := v.Unmarshal(&uluCfg); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}

	uluCfg.Viper = v
	return &uluCfg
}

func (config *ConfigBuild) Get(key string) interface{} {
	return config.Viper.Get(key)
}
