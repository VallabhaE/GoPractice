package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Information to be caried from now
// Get Request cannot hold body object which means frount-end can not send body using GET method
// but post,put,patch can access body and frount end can actually send it

//So,Learning about how to route and grouging in gin module
// the best way to learn to implement

//TASK :-1

//Complete backed for meta-verse using go lang and gin module
//for database utilise it or update it if needed
//date 4th Feb

//if i am chehing later check date how long i took to complete this backend

func main1() { //Change it to main when trying this data and change main.go func main to someothername
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			fmt.Println("v1/")
		})
		v1.POST("/login", func(ctx *gin.Context) {
			fmt.Println("v1/login")
		})
		v1.POST("/submit", func(ctx *gin.Context) {
			fmt.Println("v1/submit")
		})
	}

	router.Run(":8080")
}
