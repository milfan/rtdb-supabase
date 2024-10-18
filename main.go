package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()
	ginServer.Use(gin.Recovery())
}
