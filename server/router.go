package server

import (
	"time"
	"log/syslog"

	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/contrib/ginrus"

	"github.com/sirupsen/logrus"
	"github.com/mattn/go-colorable"
	"github.com/x-cray/logrus-prefixed-formatter"

	"github.com/gin-contrib/cors"

	"cloudfreexiao/ant-graphql/backend-go/controller"
	"cloudfreexiao/ant-graphql/backend-go/lib/user"

	lSyslog "cloudfreexiao/ant-graphql/backend-go/lib/logapi/hooks/syslog"
)

const identityKey = "cloudfreexiao"
const secretKey = "cloudfreexiao-ant-graphql"

type login struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required"`
}

type User struct {
	Name string
}

func getAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			name := loginVals.Name
			passwd := loginVals.Passwd

			if err := user.AuthenticateUser(name, passwd); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &User{
				Name: name,
			}, nil
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}

	return authMiddleware, nil
}

func buildLoggerMiddleware(on bool) (*logrus.Logger, error) {
	env := "env"

	logger := logrus.New()

	if env == "prod" {
		logger.Formatter = &logrus.JSONFormatter{}
	} else {
		formatter := &prefixed.TextFormatter{ForceColors: true, ForceFormatting: true}
		formatter.SetColorScheme(&prefixed.ColorScheme{DebugLevelStyle: "green+b", InfoLevelStyle: "green+h"})
		logger.Formatter = formatter
	}

	logger.Level = logrus.DebugLevel
	logger.Out = colorable.NewColorableStdout()

	if on {
		hook, err := lSyslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
		if err == nil {
			logger.Hooks.Add(hook)
		}
	}
	return logger, nil
}

func newRouter(c *RouterConfig) *gin.Engine {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	logMiddleware, err := buildLoggerMiddleware(!c.Debug)
	if err != nil {
		return nil
	}

	authMiddleware, err := getAuthMiddleware()
	if err != nil {
		return nil
	}

	r.Use(ginrus.Ginrus(logMiddleware, time.RFC3339, false))
	r.Use(cors.Default())
	
	ping := new(controller.PingController)

	r.GET("/ping", ping.GetController)
	r.POST("/login", authMiddleware.LoginHandler)

	r.Use(authMiddleware.MiddlewareFunc())

	graphql := new(controller.GraphqlController)

	r.POST("/refresh_token", authMiddleware.RefreshHandler)
	r.GET("/graphql", gin.WrapF(graphql.NewGraphiQLHandlerFunc()))
	r.POST("/graphql", gin.WrapH(graphql.NewHandler()))

	return r
}
