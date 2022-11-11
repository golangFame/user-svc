package router

import "github.com/gin-gonic/gin"

type Services interface {
	RoutesWithNoAuth(r *gin.RouterGroup, mws ...Middleware)
}
