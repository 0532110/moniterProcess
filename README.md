建议在linux下编译:
> go build moniterProcess.go function.go

[process]
process=chrome # 要监控的进程名
sleep=3 # 间隔时间
可以添加其他功能，如：
1、进程异常时请求指定http地址进行统计或报警处理。
2、对接钉钉机器人。
3、进程异常时自行恢复。

纯属练习～

