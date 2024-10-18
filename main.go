package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginNew := gin.Default()
	ginNew.Use(gin.Recovery())
}
