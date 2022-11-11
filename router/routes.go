package router

import (
	"github.com/BzingaApp/user-svc/handlers"
	"github.com/BzingaApp/user-svc/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//type services struct {
//	Root NoAuthRouting
//}

type service struct {
	conf         *viper.Viper
	middlewares  *middlewares.Middleware
	homeHandler  *handlers.HomeHandler
	dummyHandler *handlers.DummyHandler
}

type Middleware func() gin.HandlerFunc
