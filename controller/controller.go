package controller

import (
	"jwt-gin/dto"
	"jwt-gin/service.go"

	"github.com/gin-gonic/gin"
)

// login controller interface
type LoginController interface {
	Login(c *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func (this loginController) Login(c *gin.Context) string {
	var credential dto.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := this.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return this.jWtService.GenerateToken(credential.Email, true)
	}
	return ""
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}
