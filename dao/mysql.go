package dao

import (
	"go.mod/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql(user, pass, port, dbname string) (err error) {
	//数据库连接路径格式
	//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	//在config包中的app.json中更改配置
	dsn := user + ":" + pass + "@tcp(127.0.0.1:" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return
	}

	if err = DB.AutoMigrate(&models.Admin{}, &models.Article{}, &models.Detail{}); err != nil {
		return
	}

	return
}
