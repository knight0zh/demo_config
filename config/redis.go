package config

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisCfg struct {
	Host              string        `json:"host"`
	Port              int           `json:"port"`
	Db                int           `json:"db"`
	Password          string        `json:"password"`
	MaxIdle           int           `json:"maxIdle"`
	MaxActive         int           `json:"maxActive"`
	IdleTimeout       time.Duration `json:"idleTimeout"`
	ConnectTimeout    time.Duration `json:"connectTimeout"`
	ReadTimeout       time.Duration `json:"readTimeout"`
	WriteTimeout      time.Duration `json:"writeTimeout"`
	DefaultExpiration time.Duration `json:"defaultExpiration"`
}
type RedisClusterCfg struct {
	Common  *RedisCfg `json:"common"`
	Session *RedisCfg `json:"session"`
	Cache   *RedisCfg `json:"cache"`
}

func (c RedisCfg) NewRedis() *redis.Client {
	// 建立连接池
	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.Db,
	})

	return conn
}
