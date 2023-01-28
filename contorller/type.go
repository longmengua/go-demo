package contorller

import (
	"go-demo/service"
	"go-demo/util/server"
)

type HandlerFunc func(req interface{}) (res interface{}, err error)

type Struct struct {
	Routes  []server.Route
	Service *service.Struct
}
