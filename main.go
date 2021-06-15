package main

import (
	"go.mod/dao"
	"go.mod/routers"
	"go.mod/tool"
)

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

	//注册路由
	r := routers.SetRouter()

	err = r.Run(cfg.AppHost + ":" + cfg.HttpPort)
	//err=r.Run(":8848")
	if err != nil {
		panic(err)
	}
}
