package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Init() error {
	router := NewRouter()

	gin.SetMode(gin.ReleaseMode)

	return errors.Wrap(router.Run(":8080"), "unable  to start server")
}
