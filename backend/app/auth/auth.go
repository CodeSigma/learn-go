package auth

import (
	"fmt"
	"net/http"

	"github.com/CodeSigma/learn-go/app"
	"github.com/CodeSigma/learn-go/app/constants"
	"github.com/CodeSigma/learn-go/app/utils"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func checkAuth(c *gin.Context) {
	var auth Auth

	globalApplication := app.Gin{C: c}

	auth.Username = "user"
	auth.Password = "password"

	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Sprintf("USER %s logged in", username)

	var isAuth = auth.Username == username && auth.Password == password

	if !isAuth {
		globalApplication.Response(http.StatusInternalServerError, constants.CD_ERROR, nil)
		return
	}

	token, err := utils.GenerateToken(username, password)

	if err != nil {
		globalApplication.Response(http.StatusInternalServerError, constants.CD_ERROR, err)
		return
	}

	globalApplication.Response(
		http.StatusOK,
		constants.CD_SUCCESS,
		map[string]string{
			"token": token,
		})

}
