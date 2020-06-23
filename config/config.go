package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	DefaultConfig *ConfigBuild
)

type ConfigBuild struct {
	Viper *viper.Viper
	Mysql MysqlClusterCfg `mapstructure:"mysql"`
	Redis RedisClusterCfg `mapstructure:"redis"`
	Log   *LogCfg         `mapstructure:"log"`
}

func (config *ConfigBuild) Get(key string) interface{} {
	return config.Viper.Get(key)
}

func (config *ConfigBuild) Watch() {
	config.Viper.WatchConfig()
	config.Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		OmsMysql.Close()
		HdMysql.Close()
		CommonRedis.Close()
		InitOmsMysql()
		InitHdMysql()
		InitCommonRedis()
		if err := InitAccessLogger(config.Log); err != nil {
			log.Fatalf("init access logger failed, err:%v\n", err)
		}
		if err := InitErrorLogger(); err != nil {
			log.Fatalf("init sys logger failed, err:%v\n", err)
		}
	})
}

func InitConfig(path string) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(&DefaultConfig); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}

	DefaultConfig.Viper = v
	DefaultConfig.Watch()
}

func Get(key string) interface{} {
	return DefaultConfig.Get(key)
}

func Log() *LogCfg {
	return DefaultConfig.Log
}
