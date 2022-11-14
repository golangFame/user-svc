package handlers

import (
	"context"
	"fmt"
	"github.com/BzingaApp/user-svc/entities"
	"github.com/BzingaApp/user-svc/er"
	"github.com/BzingaApp/user-svc/services/app"
	"github.com/BzingaApp/user-svc/services/dummy"
	"github.com/BzingaApp/user-svc/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/spf13/viper"
)

type Handler struct {
	conf *viper.Viper
	log  *logrus.Logger
}

type DummyHandler struct {
	Handler

	dummyServices dummy.Services
}

type HomeHandler struct {
	Handler
}

type AppHandler struct {
	Handler

	appServices app.Services
}

func Bind(c *gin.Context, req interface{}) (statusCode int, err error) {
	if c.Request.Method == "GET" {
		queries := c.Request.URL.Query()

		if err = utils.ConvertMapToAny(queries, req); err != nil {
			c.Error(err).SetType(er.ValidationError)
			statusCode = http.StatusUnprocessableEntity
		}
		return
	}
	if err = c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		//if err = c.ShouldBind(&req); err != nil {
		c.Error(err).SetType(er.ValidationError)
		statusCode = http.StatusUnprocessableEntity
	}
	return
}

func (h *Handler) createContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(context.TODO())
	return
}

func (h *Handler) getSystemHost(c *gin.Context) (u string) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	//c.Request.Proto //can be used as well to differentiate http versions
	u = fmt.Sprintf("%s://%s", scheme, c.Request.Host)
	return
}

func (h *Handler) error(c *gin.Context, err error, errType gin.ErrorType) {
	if err != nil {
		c.Error(err).SetType(errType)
	}
}

// addErrors use with defer
func (h *Handler) addErrors(c *gin.Context, err error) {
	if err != nil {
		c.Error(err)
	}
}

func (h *Handler) handleResponse(c *gin.Context, res *entities.GenericResponse, finalError *error) {
	err := c.Errors.Last() //not working while using backup feature
	if err != nil && *finalError != nil {
		if res.StatusCode == 0 {
			res.StatusCode = http.StatusServiceUnavailable
		}
		res.Data = err
		res.Message = err.Error()
		res.Success = false
		c.AbortWithStatusJSON(res.StatusCode, res)
		return
	}
	if res.StatusCode == 0 {
		res.StatusCode = http.StatusCreated
	}
	c.JSON(res.StatusCode, res)
}
