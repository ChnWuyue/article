// @Title  router manage api
// @Description article management api by golang
// @Contact wushenxin@qq.com
package routers

import (
	"github.com/gin-gonic/gin"
	"go.mod/tool"
	"net/http"
)

func SetRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(tool.InitLogger())
	r.Use(corsMiddleware())
	r.Use(gin.Recovery())
	r.Static("/statics", "./statics")

	message(r)
	//登陆页面
	login(r)
	//管理页面
	admin(r)

	return
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
