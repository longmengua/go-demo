package main

import (
	"go-demo/contorller"
	"go-demo/entity"
	"go-demo/util/config"
	"go-demo/util/db"
	"go-demo/util/server"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		_, b, _, _ = runtime.Caller(0)
		rootPath   = filepath.Dir(b)
	)

	config := config.Struct{
		Root_Path: rootPath,
	}
	config.Init()

	db := db.Struct{
		Host:   config.DB.Host,
		Port:   config.DB.Port,
		User:   config.DB.User,
		Pwd:    config.DB.Password,
		DbName: config.DB.DB_Name,
	}
	db.Init()

	// migrate entity
	if config.DB.Migration {
		db.Instance.AutoMigrate(&entity.User{})
	}

	// declare repo

	// declare services

	// declare controllers
	controller := contorller.Struct{}
	controller.Init()

	// sum up all routes in controllers
	routes := []server.Route{}
	routes = append(routes, controller.Routes...)

	// setup api
	api := server.Struct{
		Host:   "",
		Engine: gin.Default(),
		Routes: routes,
	}
	api.Init(config.SERVER.Port)
}
