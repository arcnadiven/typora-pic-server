package main

import (
	"fmt"
	"github.com/arcnadiven/typora-pic-server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router.Init(engine)
	if err := engine.Run("0.0.0.0:8008"); err != nil {
		fmt.Println(err)
		return
	}
}
