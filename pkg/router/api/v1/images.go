package v1

import (
	"github.com/arcnadiven/typora-pic-server/pkg/service"
	"github.com/gin-gonic/gin"
)

func AddRouter(iRouter gin.IRouter) {
	apiV1Group := iRouter.Group("v1")
	{
		apiV1Group.PUT("/upload", service.Upload)
		apiV1Group.GET("/images/:dir/:name", service.Images)
	}
}
