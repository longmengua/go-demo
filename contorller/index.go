package contorller

import (
	"go-demo/contorller/user_controller"
	"go-demo/util/server"
)

func (s *Struct) Init() {
	// Inject service to controller
	userController := user_controller.Struct{
		Service: s.Service,
	}

	// Generate routes of APIs
	s.Routes = []server.Route{
		{
			Method:      "get",
			Path:        "user",
			HandlerFunc: userController.ReadUserController,
		},
		{
			Method:      "post",
			Path:        "user",
			HandlerFunc: userController.CreateUserController,
		},
		{
			Method:      "put",
			Path:        "user",
			HandlerFunc: userController.UpdateUserController,
		},
		{
			Method:      "delete",
			Path:        "user",
			HandlerFunc: userController.DeleteUserController,
		},
	}
}
