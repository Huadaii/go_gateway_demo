package main

import (
	"flag"
	"go_gateway_demo/dao"
	"go_gateway_demo/http_proxy_router"
	"go_gateway_demo/lib"
	"go_gateway_demo/router"
	"os"
	"os/signal"
	"syscall"
)

//endpoint  dashboard 后台管理 service 代理服务器
//config    ./conf/dev/  对应配置文件夹
//通过flag启动
//启动方式:目录下 go run main.go -configs=./conf/dev/ -endpoint=dashboard或server
var (
	endpoint = flag.String("endpoint", "", "input endpoint dashboard or server")
	configs  = flag.String("configs", "", "input configs file like ./conf/dev/")
)

func main() {
	//通过flag选择要开启后端还是服务端
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *configs == "" {
		flag.Usage()
		os.Exit(1)
	}

	//dashboard会启动网关后端页面。配合Vue前端或接口Swagger使用
	if *endpoint == "dashboard" {
		lib.InitModule(*configs, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		router.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		router.HttpServerStop()
	} else {
		//输入service则启动服务端功能
		lib.InitModule(*configs, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		//通过init一次性加载全部服务信息到内存
		dao.ServiceManagerHandler.LoadOnce()
		go func() {
			http_proxy_router.HttpServerRun()
		}()
		go func() {
			//https
			http_proxy_router.HttpsServerRun()
		}()
		quit := make(chan os.Signal)
		//监听收到的信号:接受ctrl+c结束程序
		//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		http_proxy_router.HttpServerStop()
		http_proxy_router.HttpsServerStop()
	}
}
