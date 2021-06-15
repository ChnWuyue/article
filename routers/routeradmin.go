package routers

import (
	"github.com/gin-gonic/gin"
	"go.mod/controllers"
)

func admin(e *gin.Engine) {
	manageController := controllers.ManageController{}

	admins := e.Group("/hub_admin")
	{
		//文章上传
		admins.POST("/article_create", controllers.CheckJWT, manageController.PutArticle)

		//动态上传
		admins.POST("/detail_create", controllers.CheckJWT, manageController.PutDetail)

		manage := admins.Group("/manage")
		{

			//返回全部文章和动态给前端
			manage.GET("/", controllers.CheckJWT, manageController.GetList)

			//文章和动态修改
			update := manage.Group("/update")
			{
				//文章
				update.GET("/article", controllers.CheckJWT, manageController.GetArticle)
				update.PUT("/article", controllers.CheckJWT, manageController.UpdateArticle)

				//动态
				update.GET("/detail", controllers.CheckJWT, manageController.GetDetail)
				update.PUT("/detail", controllers.CheckJWT, manageController.UpdateDetail)
			}

			//文章删除
			article := manage.Group("/article")
			{
				article.DELETE("/:id", controllers.CheckJWT, manageController.DeleteArticle)
			}

			//动态管理
			detail := manage.Group("/detail")
			{
				detail.PUT("/:id", controllers.CheckJWT, manageController.ShowDetail)
				detail.DELETE("/:id", controllers.CheckJWT, manageController.DeleteDetail)
			}

		}
	}
}
