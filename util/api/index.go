package api

import (
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

func (s *Struct) Init(routes []struct {
	Method      string
	Path        string
	HandlerFunc func(req interface{}) (res interface{}, err error)
}) {
	engine := gin.Default()

	// arr := &gin.RoutesInfo{}
	if routes != nil {
		for _, v := range routes {
			r := &gin.RouteInfo{
				Method: v.Method,
				Path:   v.Path,
				HandlerFunc: func(c *gin.Context) {
					responseHandler(c, v.HandlerFunc)
				},
			}
			if strings.EqualFold(r.Method, "get") {
				engine.GET(r.Path, r.HandlerFunc)
			}
			if strings.EqualFold(r.Method, "post") {
				engine.POST(r.Path, r.HandlerFunc)
			}
			if strings.EqualFold(r.Method, "put") {
				engine.PUT(r.Path, r.HandlerFunc)
			}
			if strings.EqualFold(r.Method, "delete") {
				engine.DELETE(r.Path, r.HandlerFunc)
			}
		}
	}

	s.Instance = engine
}
