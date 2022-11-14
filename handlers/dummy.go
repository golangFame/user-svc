package handlers

import (
	"github.com/BzingaApp/user-svc/entities"
	"github.com/BzingaApp/user-svc/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *DummyHandler) Dummy(c *gin.Context) {
	var (
		req = entities.Dummy{}
		res = entities.GenericResponse{}
		err error
	)

	defer h.handleResponse(c, &res, &err)

	if res.StatusCode, err = Bind(c, &req); err != nil {
		return
	}

	validator := validators.New()
	err = validator.Validating(req)
	if err != nil {
		res.StatusCode = http.StatusUnprocessableEntity
		return
	}
	h.dummyServices.Dummy(&req)

	res.Data = req
	res.Success = true
	res.Message = "Dummy services Completed"
	return
}
