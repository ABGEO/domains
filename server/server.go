package server

import (
	"github.com/pkg/errors"
)

func Init() error {
	router := NewRouter()

	return errors.Wrap(router.Run(":8080"), "unable  to start server")
}
