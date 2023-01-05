package main

import (
	"go-demo/contorller"
	"go-demo/util/api"
	"go-demo/util/config"
	"go-demo/util/db"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootPath   = filepath.Dir(b)
)

func main() {
	config := config.Struct{}
	config.Init(rootPath)

	db := db.Struct{}
	db.Init(config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password)

	controller := contorller.Struct{
		Routes: []contorller.Route{},
	}
	controller.Init(controller.Routes)

	api := api.Struct{
		Host: nil,
	}
	// todo: fix it
	// api.Init(controller.Routes)
	api.Init(nil)
}
