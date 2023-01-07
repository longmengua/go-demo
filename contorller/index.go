package contorller

import (
	"go-demo/contorller/user_controller"
	"go-demo/util/api"
)

func (s *Struct) Init() {
	s.Routes = []api.Route{
		{
			Method:      "get",
			Path:        "user",
			HandlerFunc: user_controller.ReadUserController,
		},
		{
			Method:      "post",
			Path:        "user",
			HandlerFunc: user_controller.UpsertUserController,
		},
		{
			Method:      "put",
			Path:        "user",
			HandlerFunc: user_controller.UpsertUserController,
		},
		{
			Method:      "delete",
			Path:        "user",
			HandlerFunc: user_controller.DeleteUserController,
		},
	}
}
