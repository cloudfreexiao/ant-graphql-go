package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"cloudfreexiao/ant-graphql/backend-go/lib/logapi"

	"cloudfreexiao/ant-graphql/backend-go/dao/factory"
	"cloudfreexiao/ant-graphql/backend-go/dao/mysqldb"
	"cloudfreexiao/ant-graphql/backend-go/dao/schemas"
)

type PingController struct{}

func (ctr *PingController) GetController(c *gin.Context) {
	logapi.DEBUG("++++++++", http.StatusOK)
	logapi.TRACE("++++++++", http.StatusOK)
	logapi.INFO("++++++++", http.StatusOK)
	logapi.WARN("++++++++", http.StatusOK)
	logapi.ERROR("++++++++", http.StatusOK)
	userDao, _ := factory.FactoryDao("mysql", "User")
	user := userDao.(*mysqldb.MysqlUserDao)
	schema := &schemas.User{UIN: "111", Name: "cloudfreexiao", Email: "cloudfreexiao@example.com"}
	err := user.CreateUser(schema)
	if err != nil {
		logapi.ERROR("CreateUser Error", err.Error())
	}
	u, err := user.GetUserByEmail("cloudfreexiao@example.com")
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		u.UIN = "222"
		user.UpdateUser(u)
	}
	c.String(http.StatusOK, logapi.Inspect(u))
}
