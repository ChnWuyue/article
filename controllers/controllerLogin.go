package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mod/logic"
	"go.mod/models"
	"go.mod/tool"
)

type LoginController struct {
}

// @Title [Authorize needed] Login
// @Description  login account
// @router /hubLogin [post]
func (w *LoginController) Login(c *gin.Context) {

	//接收用户数据
	var user models.Admin

	user.ID = c.PostForm("username")
	user.Password = c.PostForm("password")

	//验证

	if !(logic.AccountCheck(&user)) {
		c.JSON(200, gin.H{
			"error":   -1,
			"message": "密码或者账号输入错误",
		})
		tool.Logger.WithFields(logrus.Fields{
			"username:": user.ID,
			"password:": user.Password,
			"message:":  "密码或者账号输入错误",
		}).Info()
		return
	}

	//创建token
	tokenString, _ := NewJWT(user.ID)
	c.JSON(200, gin.H{
		"error":   0,
		"message": "success",
		"data": gin.H{
			"token": tokenString,
		},
	})

	tool.Logger.WithFields(logrus.Fields{
		"username:": user.ID,
		"password:": user.Password,
		"message:":  "账户登陆成功",
	}).Info()

	return
}
