package main

import (
	"blog/routers/defaultRouter"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Static("/static", "./static")

	engine.LoadHTMLGlob("templates/*")

	defaultRouter.Init(engine)

	engine.Run(":80")
}
