package main

import (
	"flag"
	"fmt"
	"github.com/lehoon/go-udp-server/config"
	"github.com/lehoon/go-udp-server/library/logger"
	"syscall"
)

func main() {
	flag.Parse()
	port := config.GetTcpServerPort()
	if port == 0 {
		fmt.Errorf("获取配置文件udp服务器参数失败")
		logger.Log().Error("获取配置文件udp服务器参数失败")
		return
	}

	fmt.Printf("udp.server.port为[%d]\n", port)

	c := oss.Notify(syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		if s == syscall.SIGINT || s == syscall.SIGTERM || s == syscall.SIGQUIT {
			fmt.Printf("接收到停止信号,%v\n", s)
			logger.Log().Infof("接收到停止信号,%v", s)
			break
		}
	}

	logger.Log().Info("关闭go-udp-server系统")
}