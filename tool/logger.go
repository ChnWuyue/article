package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

var Logger = logrus.New()

func init() {

	Logger.Formatter = new(logrus.TextFormatter) // default
	Logger.Formatter.(*logrus.TextFormatter).TimestampFormat = "2006年01月02日 15:04:05"
	Logger.Formatter.(*logrus.TextFormatter).FullTimestamp = true // remove timestamp from test output

	t := time.Now()
	now := t.Format("2006年01月02日")

	timeStr := fmt.Sprintf("%s/%s/", now[:7], now[7:12])

	logPath := "./log/backend/" + timeStr
	ok, _ := IsFileExist(logPath)
	if !ok {
		os.MkdirAll(logPath, os.ModePerm)
	}
	logName := fmt.Sprintf("%d%s", time.Now().Day(), "号日志.log")
	src, err := os.OpenFile(logPath+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("err", err)
	}
	Logger.Out = src
}

func InitLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		//花费时间
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		//请求主机
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		//请求状态码
		statusCode := c.Writer.Status()
		//客户端ip
		clientIp := c.ClientIP()
		//客户端信息
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		//请求的方法
		method := c.Request.Method
		//请求的uri
		path := c.Request.RequestURI

		entry := Logger.WithFields(logrus.Fields{
			"hostName":  hostName,
			"status":    statusCode,
			"spendTime": spendTime,
			"ip":        clientIp,
			"method":    method,
			"path":      path,
			"dataSize":  dataSize,
			"agent":     userAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode > 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
