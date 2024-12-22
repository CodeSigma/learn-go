package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to go-tweet API",
		})
	})

	return r
}
