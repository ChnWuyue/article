package main

import (
	"fmt"
	"go.mod/dao"
	"go.mod/models"
	"go.mod/tool"
)

/*import (
	"go.mod/logic"
	"go.mod/models"
	"testing"
)

func TestAddArticle(t *testing.T) {
	var article *models.Article
	article = &models.Article{
		ID: "123456",
	}
	logic.AddArticle(article)
}
*/

func main() {
	//配置
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	//数据库
	err = dao.InitMysql(cfg.DBUser, cfg.DBPass, cfg.DBPort, cfg.DBName)
	if err != nil {
		panic(err)
	}
	fmt.Println(GetRelevantDetail("5"))
}

func GetRelevantDetail(id string) []models.Detail {
	var details []models.Detail
	var previousDetail models.Detail
	var nextDetail models.Detail
	dao.DB.Where("id < ?", id).Last(&previousDetail)
	dao.DB.Where("id > ?", id).First(&nextDetail)
	details = append(details, previousDetail, nextDetail)
	return details
}
