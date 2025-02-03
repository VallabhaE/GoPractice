package funcs


import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func INIT(r *gin.Engine) {
	r.Use(func(ctx *gin.Context) {
		fmt.Println("MIDDLE WERE 1")
		ctx.Next()
	})

	r.Use(func(ctx *gin.Context) {
		fmt.Println("MIDDLE WERE 2")
		ctx.Next()
	})
}