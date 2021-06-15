package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mod/logic"
	"go.mod/models"
	"net/http"
)

type IndexController struct {
}

// @Title [Authorize needed] IndexDisplay
// @Description  返回主页需要显示的信息
// @router /index [get]
func (w *IndexController) IndexDisplay(c *gin.Context) {
	var detail []models.Detail
	var details []models.Detail
	detail = logic.DisplayDetail()
	length := len(detail)
	if length < 3 {
		details = detail
	} else {
		details = detail[0:3]
	}

	var article []models.Article
	var articles []models.Article
	article = logic.GetArticleList()
	length = len(article)
	if length < 5 {
		articles = article
	} else {
		articles = article[0:5]
	}

	c.JSON(http.StatusOK, gin.H{
		"detail":  details,
		"article": articles,
	})

}

func (w *IndexController) DetailDisplay(c *gin.Context) {
	id := c.Query("id")
	detail := logic.GetDetailById(id)
	details := logic.GetRelevantDetail(id)
	previousDetail := details[0]
	nextDetail := details[1]
	c.JSON(http.StatusOK, gin.H{
		"detail": map[string]interface{}{
			"id":          detail.ID,
			"title":       detail.Title,
			"abstract":    detail.Abstract,
			"create_time": detail.CreateTime,
			"html":        detail.Html,
			"url":         detail.Cover,
		},
		"otherDetail": map[string]interface{}{
			"previous": map[string]interface{}{
				"id":          previousDetail.ID,
				"title":       previousDetail.Title,
				"abstract":    previousDetail.Abstract,
				"create_time": previousDetail.CreateTime,
			},
			"next": map[string]interface{}{
				"id":          nextDetail.ID,
				"title":       nextDetail.Title,
				"abstract":    nextDetail.Abstract,
				"create_time": nextDetail.CreateTime,
			},
		},
	})
}

func (w *IndexController) ArticleDisplay(c *gin.Context) {
	id := c.Query("id")
	article := logic.GetArticleById(id)
	var send string
	if article.File == "" {
		send = article.Link
	} else {
		send = article.File
	}
	c.JSON(http.StatusOK, gin.H{
		"article": map[string]interface{}{
			"id":          article.ID,
			"title":       article.Title,
			"author":      article.Author,
			"create_time": article.CreateTime,
			"html":        article.Html,
			"url":         send,
		},
	})
}

/*// @Title [Authorize needed] ParseIndex
  // @Description  Parse web
  // @router /index [post]
  func (w *IndexController) ParseIndex(c *gin.Context) {
  	c.HTML(http.StatusOK, "index.html", gin.H{})
  	return
  }

  func (w *IndexController) JumpDetails(c *gin.Context) {
  	c.Redirect(http.StatusFound, "/details")
  	return
  }
*/
