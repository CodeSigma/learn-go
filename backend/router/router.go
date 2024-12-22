package router

import (
	"github.com/CodeSigma/learn-go/app/auth"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to go-tweet API",
		})
	})

	r.GET("/auth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Doing Authentication",
		})
	})

	apiv1 := r.Group("/api/v1")
	apiv1.Use(auth.JWT())
	{
		apiv1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Private API",
			})
		})
	}

	return r
}
