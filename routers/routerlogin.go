package routers

import (
	"github.com/gin-gonic/gin"
	"go.mod/controllers"
)

func login(e *gin.Engine) {
	loginController := controllers.LoginController{}

	e.POST("/hubLogin", loginController.Login)

}
