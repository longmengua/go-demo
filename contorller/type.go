package contorller

import "go-demo/util/api"

type HandlerFunc func(req interface{}) (res interface{}, err error)

type Struct struct {
	Routes []api.Route
}
