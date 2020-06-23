package config

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	OmsMysql *gorm.DB
	HdMysql  *gorm.DB
)

type MysqlCfg struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	MaxIdle  int    `json:"maxIdle"`
	MaxOpen  int    `json:"maxOpen"`
	Debug    bool   `json:"debug"`
}

type MysqlClusterCfg struct {
	HuiDian *MysqlCfg `mapstructure:"db_huidian"`
	OMS     *MysqlCfg `mapstructure:"db_oms"`
}

func (c *MysqlCfg) NewMysql() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@(%s:%d)/?charset=utf8&parseTime=true&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
	)

	db, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Fatalln(err)
	}
	if c.Debug {
		db.LogMode(c.Debug)
	}
	db.DB().SetMaxIdleConns(c.MaxIdle)
	db.DB().SetMaxOpenConns(c.MaxOpen)
	db.DB().SetConnMaxLifetime(time.Second * 300)

	db.SingularTable(true)
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	return db
}

func InitOmsMysql() {
	OmsMysql = DefaultConfig.Mysql.OMS.NewMysql()
}

func InitHdMysql() {
	HdMysql = DefaultConfig.Mysql.OMS.NewMysql()
}
