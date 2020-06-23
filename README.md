# demo-config

## 介绍
配置管理包用于初始化`db`,`redis`,`log`,`config`,惠店业务的分表逻辑。
如需增加`db`,`redis`实例需要在包中加入。

## 目录结构
    ├── README.md
    ├── config
    │   ├── config.go
    │   ├── log.go
    │   ├── mysql.go
    │   └── redis.go
    ├── go.mod
    ├── go.sum
    └── spilt
        ├── split.go
        └── split_map.go
