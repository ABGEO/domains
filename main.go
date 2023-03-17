package main

import (
	"github.com/abgeo/domains/config"
	"github.com/abgeo/domains/server"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	if err := server.Init(); err != nil {
		panic(err)
	}
}
