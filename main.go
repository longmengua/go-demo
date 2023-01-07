package main

import (
	"go-demo/contorller"
	"go-demo/entity"
	"go-demo/util/api"
	"go-demo/util/config"
	"go-demo/util/db"
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
		RootPath: rootPath,
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
	if config.MIGRATE_UP {
		db.Instance.AutoMigrate(&entity.User{})
	}

	// declare controllers
	controller := contorller.Struct{}
	controller.Init()

	// sum up all routes in controllers
	routes := []api.Route{}
	routes = append(routes, controller.Routes...)

	// setup api
	api := api.Struct{
		Host:   "",
		Engine: gin.Default(),
		Routes: routes,
	}
	api.Init()
}
