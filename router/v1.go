package router

import (
	"github.com/arcnadiven/typora-pic-server/service"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	v1Grp := engine.Group("v1")
	{
		v1Grp.PUT("/upload", service.Upload)
		v1Grp.GET("/images/:name", service.Images)
	}
}
