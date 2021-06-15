package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mod/logic"
	"go.mod/models"
	"go.mod/tool"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type ManageController struct {
}

func getArticle(article *models.Article, c *gin.Context) {
	article.Title = c.PostForm("title")
	article.CreateTime = c.PostForm("create_time")
	article.UpdateTime = c.PostForm("update_time")
	article.Html = c.PostForm("html")
	article.Author = c.PostForm("author")
	article.Link = c.PostForm("link")
}

func getDetail(detail *models.Detail, c *gin.Context) {
	detail.Title = c.PostForm("title")
	detail.CreateTime = c.PostForm("create_time")
	detail.UpdateTime = c.PostForm("update_time")
	detail.Html = c.PostForm("html")
	detail.Abstract = c.PostForm("abstract")
}

// @Title [Authorize needed] PutArticle
// @Description  上传文章
// @router /hubAdmin/articleCreate [post]
func (w *ManageController) PutArticle(c *gin.Context) {

	var errorStr string
	article := &models.Article{}
	getArticle(article, c)

	article.File, errorStr = uploadArticleFile(c)
	if article.Link == "" && errorStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   -4,
			"message": errorStr,
		})
		tool.Logger.WithFields(logrus.Fields{
			"文章id":    article.ID,
			"message": "上传失败，文件或链接问题",
		}).Info()
		return
	}

	logic.AddArticle(article)

	c.JSON(http.StatusCreated, gin.H{
		"error":   0,
		"message": "上传成功",
	})
	tool.Logger.WithFields(logrus.Fields{
		"文章id":    article.ID,
		"message": "上传成功",
	}).Info()
	return
}

// @Title [Authorize needed] UpdateArticle
// @Description  更新文章
// @router /hubAdmin/manage/article [put]
func (w *ManageController) UpdateArticle(c *gin.Context) {

	id := c.PostForm("id")
	article := &models.Article{}
	article.ID, _ = strconv.Atoi(id)

	errorStr := ""
	getArticle(article, c)
	if file := c.PostForm("file"); file == "" {
		article.File, errorStr = uploadArticleFile(c)
	} else {
		article.File = file
	}

	if article.Link == "" && errorStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   -4,
			"message": errorStr,
		})
		tool.Logger.WithFields(logrus.Fields{
			"文章id":    article.ID,
			"message": "修改失败，链接或文件问题",
		}).Info()
		return
	}

	logic.UpdateArticle(article)

	c.JSON(http.StatusCreated, gin.H{
		"err":     0,
		"message": "修改文章成功",
	})
	tool.Logger.WithFields(logrus.Fields{
		"文章id":    article.ID,
		"message": "上传成功",
	}).Info()
	return
}

// @Title [Authorize needed] GetArticle
// @Description  通过id得到动态信息
// @router /hubAdmin/manage/update/article?id= [GET]
func (w *ManageController) GetArticle(c *gin.Context) {
	id := c.Query("id")
	article := logic.GetArticleById(id)

	c.JSON(200, gin.H{
		"err":     0,
		"message": "修改的文章的id为" + id,
		"article": article,
	})

	tool.Logger.WithFields(logrus.Fields{
		"被修改文章的id": id,
	})

	return
}

// @Title [Authorize needed] GetDetail
// @Description  通过id得到动态信息
// @router /hubAdmin/manage/update/detail?id= [GET]
func (w *ManageController) GetDetail(c *gin.Context) {
	id := c.Query("id")
	detail := logic.GetDetailById(id)

	c.JSON(200, gin.H{
		"err":     0,
		"message": "修改的文章的id为" + id,
		"detail":  detail,
	})

	tool.Logger.WithFields(logrus.Fields{
		"被修改动态的id": id,
	})
	return
}

// @Title [Authorize needed] DeletedArticle
// @Description  删除文章
// @router /hubAdmin/manage/article [delete]
func (w *ManageController) DeleteArticle(c *gin.Context) {
	var ID string
	ID = c.Param("id")
	logic.DeleteArticle(ID)

	c.JSON(http.StatusCreated, gin.H{
		"err":     0,
		"message": "文章删除成功",
	})
	tool.Logger.WithFields(logrus.Fields{
		"文章id":    ID,
		"message": "删除成功",
	}).Info()
	return
}

// @Title [Authorize needed] GetList
// @Description  范围全部文章和动态
// @router /hubAdmin/manage [get]
func (w *ManageController) GetList(c *gin.Context) {
	var details []models.Detail
	var articles []models.Article

	details = logic.GetDetail()
	articles = logic.GetArticleList()

	c.JSON(http.StatusOK, gin.H{
		"article": articles,
		"detail":  details,
	})
	return
}

// @Title [Authorize needed] ShowDetail
// @Description  选择是否显示动态
// @router /hub_admin/manage/detail/:id [put]
func (w *ManageController) ShowDetail(c *gin.Context) {
	id := c.Param("id")
	display := c.PostForm("display")
	if display != "0" && display != "1" {
		c.JSON(200, gin.H{
			"err":     -6,
			"message": "传入错误",
		})
		tool.Logger.WithFields(logrus.Fields{
			"动态id":    id,
			"message": "修改显示失败传参不为0或1",
		}).Info()
		return
	}
	logic.IsDisplay(id, display)
	if display == "1" {
		c.JSON(200, gin.H{
			"err":     0,
			"message": "修改成功现在为显示状态",
		})
		tool.Logger.WithFields(logrus.Fields{
			"动态id":    id,
			"message": "修改为显示状态",
		}).Info()
	} else if display == "0" {
		c.JSON(200, gin.H{
			"err":     0,
			"message": "修改成功现在为不显示状态",
		})
		tool.Logger.WithFields(logrus.Fields{
			"动态id":    id,
			"message": "修改为不显示状态",
		}).Info()
	}
	return
}

// @Title [Authorize needed] PutDetail
// @Description  上传动态
// @router /hubAdmin/articleCreate [post]
func (w *ManageController) PutDetail(c *gin.Context) {

	var errorStr string
	detail := &models.Detail{}
	getDetail(detail, c)
	detail.Cover, errorStr = uploadImg(c)
	if errorStr != "" {
		c.JSON(http.StatusCreated, gin.H{
			"error":   -5,
			"message": errorStr,
		})
		tool.Logger.WithFields(logrus.Fields{
			"动态id":    detail.ID,
			"message": "上传失败，图片问题",
		}).Info()
		return
	}

	logic.AddDetail(detail)

	c.JSON(http.StatusCreated, gin.H{
		"message": "上传成功",
	})
	tool.Logger.WithFields(logrus.Fields{
		"动态id":    detail.ID,
		"message": "上传成功",
	}).Info()
	return
}

// @Title [Authorize needed] UpdateDetail
// @Description  更新动态
// @router /hubAdmin/manage/article [put]
func (w *ManageController) UpdateDetail(c *gin.Context) {

	id := c.PostForm("id")
	detail := &models.Detail{}
	detail.ID, _ = strconv.Atoi(id)

	errorStr := ""
	getDetail(detail, c)
	if cover := c.PostForm("cover"); cover == "" {
		detail.Cover, errorStr = uploadImg(c)
	} else {
		detail.Cover = cover
	}
	if errorStr != "" {
		c.JSON(http.StatusCreated, gin.H{
			"error":   -5,
			"message": errorStr,
		})
		tool.Logger.WithFields(logrus.Fields{
			"动态id":    detail.ID,
			"message": "修改失败，图片问题",
		}).Info()
		return
	}
	logic.UpdateDetail(detail)

	c.JSON(http.StatusCreated, gin.H{
		"err":     0,
		"message": "修改动态成功",
	})
	tool.Logger.WithFields(logrus.Fields{
		"动态id":    detail.ID,
		"message": "修改成功",
	}).Info()
	return
}

// @Title [Authorize needed] DeleteDetail
// @Description  删除动态
// @router /hubAdmin/manage/detail [delete]
func (w *ManageController) DeleteDetail(c *gin.Context) {
	var ID string
	ID = c.Param("id")
	logic.DeleteDetail(ID)
	c.JSON(http.StatusCreated, gin.H{
		"err":     0,
		"message": "动态删除成功",
	})
	tool.Logger.WithFields(logrus.Fields{
		"动态id":    ID,
		"message": "删除成功",
	}).Info()
	return
}

func uploadImg(c *gin.Context) (string, string) {

	f, err := c.FormFile("cover")
	if err != nil {
		errorStr := "上传失败，可能为文件过大，或者没有选择图片"
		return "", errorStr
	} else {
		fileExt := strings.ToLower(path.Ext(f.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			errorStr := "上传失败文件只能为.jpg .gif .jpeg 格式"
			return "", errorStr
		}

		imgName := fmt.Sprint(time.Now().Day(), "号", time.Now().Hour(), "点封面", filepath.Base(f.Filename))
		fmt.Println(imgName)
		imgDir := fmt.Sprintf("%s%d%s/", "./statics/cover/", time.Now().Year(), time.Now().Month().String())
		fmt.Println(imgDir)
		isExist, _ := tool.IsFileExist(imgDir)
		if !isExist {
			err = os.MkdirAll(imgDir, os.ModePerm)
		}

		imgPath := fmt.Sprintf("%s%s", imgDir, imgName)
		err = c.SaveUploadedFile(f, imgPath)
		if err != nil {
			errorStr := "下载到本地失败"
			return "", errorStr
		}

		return imgPath[1:], ""
	}
}

func uploadArticleFile(c *gin.Context) (string, string) {

	f, err := c.FormFile("file")
	if err != nil {
		errorStr := "上传失败"
		return "", errorStr
	} else {
		fileExt := strings.ToLower(path.Ext(f.Filename))
		if fileExt != ".txt" && fileExt != ".md" && fileExt != ".doc" && fileExt != ".docx" && fileExt != ".pdf" {
			errorStr := "上传失败文件不是指定的文本文件"
			return "", errorStr
		}

		fileName := fmt.Sprint(time.Now().Day(), "号", time.Now().Hour(), "点", filepath.Base(f.Filename))
		fileDir := fmt.Sprintf("%s%d%s/", "./statics/article/", time.Now().Year(), time.Now().Month().String())

		isExist, _ := tool.IsFileExist(fileDir)
		if !isExist {
			err = os.MkdirAll(fileDir, os.ModePerm)
		}

		filePath := fmt.Sprintf("%s%s", fileDir, fileName)
		err = c.SaveUploadedFile(f, filePath)
		if err != nil {
			errorStr := "下载到本地失败"
			fmt.Println(err)
			return "", errorStr
		}

		return filePath[1:], ""
	}
}

func (w *ManageController) GetTime(c *gin.Context) {

	t := time.Now()
	now := t.Format("2006年1月2日")

	c.JSON(200, gin.H{
		"time": now,
	})
}
