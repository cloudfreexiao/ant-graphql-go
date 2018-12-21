package dao

import (
	"cloudfreexiao/ant-graphql/backend-go/dao/factory"
	"cloudfreexiao/ant-graphql/backend-go/dao/mysqldb"
)

func init() {
	factory.Init()
	mysqldb.Init()
}
