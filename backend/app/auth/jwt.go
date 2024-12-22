package auth

import (
	"net/http"

	"github.com/CodeSigma/learn-go/app/constants"
	"github.com/CodeSigma/learn-go/app/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = constants.CD_SUCCESS
		token := c.Query("token")
		if token == "" {
			code = constants.CD_INVALID_PARAMS
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = constants.CD_ERROR
				default:
					code = constants.CD_ERROR
				}
			}
		}

		if code != constants.CD_SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  constants.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
