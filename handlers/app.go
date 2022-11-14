package handlers

import (
	"github.com/BzingaApp/user-svc/entities"
	"github.com/BzingaApp/user-svc/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *AppHandler) Home(c *gin.Context) {
	var (
		res = entities.GenericResponse{}
		err error
	)

	defer h.handleResponse(c, &res, &err)

	res.Success = true
	res.Message = "Home service is up and running"
}

func (h *AppHandler) HomePage(c *gin.Context) {

	var (
		res entities.GenericResponse
		err error
	)
	defer h.handleResponse(c, &res, &err)

	userID := utils.ConvertStringIntoInt(c.Param("userID"))

	if userID == 0 {
		err = errors.New("invalid user")
	}

}
