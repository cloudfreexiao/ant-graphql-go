package factory

import (
	"fmt"
	"reflect"

	"cloudfreexiao/ant-graphql/backend-go/dao/mysqldb"

	"cloudfreexiao/ant-graphql/backend-go/lib/logapi"
)

//定义注册结构map
type SchemasStructMap struct {
	maps map[string]reflect.Type
}

var schemasMap *SchemasStructMap = &SchemasStructMap{make(map[string]reflect.Type)}

//根据名字注册实例
func (ssm *SchemasStructMap) Register(name string, c interface{}) {
	ssm.maps[name] = reflect.TypeOf(c).Elem()
}

//根据name初始化结构
//在这里根据结构的成员注解进行DI注入，这里没有实现，只是简单都初始化
func (ssm *SchemasStructMap) NewSchema(name string) (interface{}, error) {
	var c interface{}
	var err error

	if v, ok := ssm.maps[name]; ok {
		c = reflect.New(v).Interface()
		logapi.DEBUG("schema found ", name, reflect.TypeOf(c))
		return c, nil
	} else {
		logapi.ERROR("schema not found", name, len(ssm.maps))
		err = fmt.Errorf("schema not found %s struct", name)
	}
	return nil, err
}

func Init() {
	initMysqlSchemas()
}

func FactoryDao(drivername string, tbname string) (interface{}, error) {
	var d interface{}
	var e error

	switch drivername {
	case "mysql":
		d, e = schemasMap.NewSchema(tbname)
	default:
		logapi.ERROR("El motor [%s]:[%s] no esta implementado", drivername, tbname)
		e = fmt.Errorf("driver not found %s", drivername)
		return nil, e
	}
	return d, e
}

func initMysqlSchemas() {
	schemasMap.Register("User", &mysqldb.MysqlUserDao{})
	logapi.DEBUG("init schemasMap ", schemasMap)
}
