package router

import (
	"github.com/gin-gonic/gin"
)

func (s *service) RoutesWithNoAuth(r *gin.RouterGroup, mws ...Middleware) {
	//r := router.Group("/")

	for _, i := range mws {
		r.Use(i())
	}

	r.GET("/", s.homeHandler.Home)

	r.POST("/dummy", s.dummyHandler.Dummy)

	appRouter := r.Group("/app")
	appRouter.GET("/", s.appHandler.Home)

}
