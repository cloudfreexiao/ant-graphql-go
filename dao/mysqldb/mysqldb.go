package mysqldb

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	"cloudfreexiao/ant-graphql/backend-go/dao/schemas"
	logapi "cloudfreexiao/ant-graphql/backend-go/lib/logapi"
)

var engine *xorm.Engine

const configPrefix = "config/db/"

func Init() {
	initEngine()
	initTables()
}

func initEngine() {
	var err error
	//读取配置文件
	cfg, err := ini.Load(configPrefix + "db.ini")
	if err != nil {
		logapi.ERROR("load db config error:", err.Error())
	}
	username := cfg.Section("mysql").Key("username").Value()
	password := cfg.Section("mysql").Key("password").Value()
	url := cfg.Section("mysql").Key("url").Value()

	// 初始化xorm的engine
	source := fmt.Sprintf("%s:%s@%s", username, password, url)
	engine, err = xorm.NewEngine("mysql", source)
	if err != nil {
		logapi.ERROR("new db engine error:", err.Error())
	}

	err = engine.Ping()
	if err != nil {
		logapi.ERROR("db connect error:", err.Error())
	}

	// 30minute ping db to keep connection
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				logapi.ERROR("db connect error:", err.Error())
			}
		}
	}(engine)

	//答应日志
	engine.ShowSQL(true)
	//设置连接池大小
	engine.SetMaxIdleConns(4)
	engine.SetMaxOpenConns(8)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.SnakeMapper{})
}

func initTables() {
	engine.Sync2(
		new(schemas.User),
	)
}

func get() *xorm.Engine {
	return engine
}
