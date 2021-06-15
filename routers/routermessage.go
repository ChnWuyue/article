package routers

import (
	"github.com/gin-gonic/gin"
	"go.mod/controllers"
)

func message(e *gin.Engine) {
	indexController := controllers.IndexController{}
	e.GET("/", indexController.IndexDisplay)
	e.GET("/index", indexController.IndexDisplay)
	e.GET("/detail", indexController.DetailDisplay)
	e.GET("/article", indexController.ArticleDisplay)
}
