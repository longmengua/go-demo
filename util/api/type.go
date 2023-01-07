package api

import "github.com/gin-gonic/gin"

type HandlerFunc func(req interface{}) (res interface{}, err error)
type Route struct {
	Method      string
	Path        string
	HandlerFunc HandlerFunc
}
type Struct struct {
	Host   string
	Engine *gin.Engine
	Routes []Route
}
