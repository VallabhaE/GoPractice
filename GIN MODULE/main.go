package main

import (
	"fmt"
	"io"
	"main/funcs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	funcs.INIT(r)
	r.GET("/", func(c *gin.Context) {	
		fmt.Println("THIS IS GET REQUEST",c.Request.Body)
		body, _ := io.ReadAll(c.Request.Body)
		fmt.Println(string(body))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/", func(c *gin.Context) {	
		body, _ := io.ReadAll(c.Request.Body)
		fmt.Println(string(body))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	
	r.Run() // listen and serve on 0.0.0.0:8080
}
