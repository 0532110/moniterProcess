package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main()  {

	// 检查进程是否存在
	for true {
		processName := getConfig("process")
		checkProcess := getProcessStatus(processName)

		if checkProcess == "1" {
			fmt.Println(getTime() + " " + processName + " " + "进程运行中")
		}else if checkProcess == "0" {
			fmt.Println(getTime() + " " + processName + " " + "进程不存在")
		}else {
			fmt.Println(checkProcess)
		}

		var sleepTime int
		sleepTime = str2int(getConfig("sleep"))

		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}

func getConfig(key string) string {
	cfg,err := goconfig.LoadConfigFile("./config.ini")

	if err != nil {
		return "读取错误： " + err.Error()
	}

	value,err := cfg.GetValue("process", key)

	if err != nil {
		return "获取值失败： " + err.Error()
	}

	return value
}

func getProcessStatus(processName string) string {
	cmd := exec.Command("ps", "aux")
	rs,err := cmd.Output()

	if err != nil {
		return "执行失败： " + err.Error()
	}

	check := strings.Contains(string(rs), processName)

	if check {
		return "1"
	}else {
		return "0"
	}
}

func str2int(str string) int {
	rs,err := strconv.Atoi(str)
	if err == nil {
		return rs
	}
	return rs
}