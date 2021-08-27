package main

import (
	"jwt-gin/controller"
	"jwt-gin/service.go"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jWtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jWtService)

	app := gin.Default()

	app.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Failed to login",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	app.Run(":8080")
}
