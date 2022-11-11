package er

import "github.com/gin-gonic/gin"

const (
	ValidationError gin.ErrorType = 1
	BadRequest      gin.ErrorType = 2
)

const (
	UnknownError    gin.ErrorType = 0
	ArtificialError gin.ErrorType = 1001
	SystemError     gin.ErrorType = 1002
	ServiceError    gin.ErrorType = 1003
)
