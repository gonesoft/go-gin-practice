package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	//This is an endpoint
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OKAY!",
		})
	})

	//initialize the server
	server.Run(":8080")
}
