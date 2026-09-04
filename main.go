package main

import (
	"flag"
	"github.com/arcnadiven/typora-pic-server/pkg/router/api"
	"github.com/arcnadiven/typora-pic-server/pkg/syncer"
	"github.com/gin-gonic/gin"
)

var (
	port = flag.String("port", "8008", "Port to listen on")
)

func main() {
	flag.Parse()

	syncer.Run()
	
	engine := gin.Default()
	api.AddRouter(engine)

	if err := engine.Run(":" + *port); err != nil {
		panic(err)
	}
}
