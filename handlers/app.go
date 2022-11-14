package handlers

import (
	"github.com/BzingaApp/user-svc/entities"
	"github.com/gin-gonic/gin"
)

func (h *AppHandler) Home(c *gin.Context) {
	var (
		req = entities.Dummy{}
		res = entities.GenericResponse{}
		err error
	)

	//ctx, cancel := h.createContext()
	//defer cancel()
	//go h.dataDogTracer(ctx, &err)

	defer h.handleResponse(c, &res, &err)

	res.Data = req
	res.Success = true
	res.Message = "Home service is up and running"
	return
}
