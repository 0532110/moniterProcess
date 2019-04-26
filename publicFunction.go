package main

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// 封装get post请求，支持添加自定义header
// 用法： url 地址, method 方法 ,header 头信息 , body   数据
func GetPost(url string, method string, header map[string]string, body string) string {

	c := &http.Client{}
	r,e := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(body))

	if e != nil {
		return "初始化失败，原因：" + e.Error()
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 注入自定义头信息 - 开始
	for k, v := range header {
		r.Header.Add(k,v)  // 头部会自动添加 HTTP_ 字符
	}
	// 注入自定义头信息 - 结束

	rs,_ := c.Do(r)
	defer rs.Body.Close()
	b, err := ioutil.ReadAll(rs.Body)

	if err != nil {
		return "请求失败，原因：" + err.Error()
	}else {
		return string(b)
	}
}

func getTime() string {
	var currentTime string
	currentTime = time.Now().Format("2006-01-02 15:04:05") // 直接记为 12345 3可以表示03,12小时制，15为24小时制
	//fmt.Println(reflect.TypeOf(currentTime)) // 打印数据类型
	return currentTime
}

func writeLog(content string) {
	// 日志名
	file := time.Now().Format("20060102") + ".txt" //文件名
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	log.SetOutput(logFile)

	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:true})

	log.Info(content)
	//log.Error("error message")
	//log.Warn("warnning message")
	//log.Debug("debug message")
	//log.Fatal("fatal message")
}