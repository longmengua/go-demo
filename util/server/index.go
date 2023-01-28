package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func responseHandler(ctx *gin.Context, h func(req interface{}) (res interface{}, err error)) {
	res, err := h(ctx.Request)

	if err != nil {
		// todo: Format error
		ctx.String(http.StatusBadRequest, "Error")
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Struct) Init(port int) {
	if s.Routes != nil {
		for _, v := range s.Routes {
			r := &gin.RouteInfo{
				Method: v.Method,
				Path:   v.Path,
				HandlerFunc: func(c *gin.Context) {
					responseHandler(c, v.HandlerFunc)
				},
			}
			if strings.EqualFold(r.Method, "get") {
				s.Engine.GET(r.Path, r.HandlerFunc)
			}
			if strings.EqualFold(r.Method, "post") {
				s.Engine.POST(r.Path, r.HandlerFunc)
			}
			if strings.EqualFold(r.Method, "put") {
				s.Engine.PUT(r.Path, r.HandlerFunc)
			}
			if strings.EqualFold(r.Method, "delete") {
				s.Engine.DELETE(r.Path, r.HandlerFunc)
			}
		}
	}

	s.Engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Server alive")
	})

	s.Engine.SetTrustedProxies([]string{"localhost"})

	// launch api server
	portStr := fmt.Sprintf(":%d", port)
	s.Engine.Run(portStr)
}
