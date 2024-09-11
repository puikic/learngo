package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello,this is 参数路由！"+name)
	})
	r.GET("/order", func(ctx *gin.Context) {
		oid := ctx.Query("id")
		ctx.String(http.StatusOK, "hello 这是查询参数"+oid)
	})
	r.GET("/items/*abc", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello ,这是items")
	})
	r.Run(":8080")
}
