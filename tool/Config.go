// @Title  config tool
// @Description  配置
// @Contact wushenxin@qq.com
package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName    string `json:"app_name"`
	HttpPort   string `json:"http_port"`
	AppHost    string `json:"app_host"`
	RunMode    string `json:"run_mode"`
	AutoRender string `json:"auto_render"`
	DBUser     string `json:"database_user"`
	DBPass     string `json:"database_password"`
	DBPort     string `json:"database_port"`
	DBName     string `json:"database_dbname"`
}

var _cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&_cfg)
	if err != nil {
		return nil, err
	}
	return _cfg, nil
}
