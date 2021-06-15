package logic

import (
	"go.mod/dao"
	"go.mod/models"
	"go.mod/tool"
)

// @Title [Authorize needed] AddArticle
// @Description  Add Article
func AddArticle(article *models.Article) {
	dao.DB.Create(article)
}

// @Title [Authorize needed] UpdateArticle
// @Description  Update Article
func UpdateArticle(article *models.Article) {
	articleMap := map[string]interface{}{
		"id":          article.ID,
		"title":       article.Title,
		"create_time": article.CreateTime,
		"update_time": article.UpdateTime,
		"html":        article.Html,
		"author":      article.Author,
		"views":       article.Views,
		"file":        article.File,
		"link":        article.Link,
	}
	dao.DB.Model(&article).Updates(articleMap)
}

// @Title [Authorize needed] DeleteArticle
// @Description  Deleted Article
func DeleteArticle(ID string) {
	article := models.Article{}
	dao.DB.First(&article, ID)
	exist, _ := tool.IsFileExist(article.File)
	if exist {
		_ = tool.RemoveFile(article.File)
	}
	dao.DB.Delete(&models.Article{}, ID)
}

// @Title [Authorize needed] GetArticleList
// @Description  返回全部文章
func GetArticleList() []models.Article {
	var article []models.Article
	dao.DB.Find(&article)
	return article
}

// @Title [Authorize needed] GetArticleById
// @Description  通过id查找文章
func GetArticleById(id string) models.Article {
	article := models.Article{}
	dao.DB.First(&article, id)
	return article
}

// @Title [Authorize needed] FindArticleByID
// @Description  Find Article By ID
// @router /admin [post]
/*func FindArticleByID(ID string) models.Article {
	return
}*/
