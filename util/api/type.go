package api

import "github.com/gin-gonic/gin"

type Struct struct {
	Host     *string
	Instance *gin.Engine
}
