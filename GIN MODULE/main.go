package main

import (
	"github.com/gin-gonic/gin"
	"main/funcs"
)

func main() {
	r := gin.Default()
	funcs.INIT(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	
	r.Run() // listen and serve on 0.0.0.0:8080
}
