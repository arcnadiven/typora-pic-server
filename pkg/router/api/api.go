package api

import (
	"github.com/arcnadiven/typora-pic-server/pkg/router/api/v1"
	"github.com/gin-gonic/gin"
)

func AddRouter(iRouter gin.IRouter) {
	apiGroup := iRouter.Group("api")

	v1.AddRouter(apiGroup)
}
