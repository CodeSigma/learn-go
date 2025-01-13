package router

import (
	"github.com/CodeSigma/learn-go/app/auth"
	apiConstant "github.com/CodeSigma/learn-go/app/constants/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// Define routes
	r.GET(apiConstant.HOME, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to go-tweet API",
		})
	})

	r.GET(apiConstant.AUTH, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Doing Authentication",
		})
	})

	apiv1 := r.Group(apiConstant.GROUP_APIV1)
	apiv1.Use(auth.JWT())
	{
		apiv1.GET(apiConstant.TEST, func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Private API",
			})
		})
	}

	return r
}
