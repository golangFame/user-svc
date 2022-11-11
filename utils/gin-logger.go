package utils

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
