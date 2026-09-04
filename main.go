package main

import (
	"github.com/arcnadiven/typora-pic-server/pkg/router/api"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	api.AddRouter(engine)
	if err := engine.Run(":8008"); err != nil {
		panic(err)
	}
}
