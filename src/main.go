package main

import (
	"log"
	"net/http"
)

func main() {
	//读取配置文件
	addr := myServer.App.IP + ":" + myServer.App.Port

	//静态资源路由注册
	staticMux.RegisterStaticResource()

	err := http.ListenAndServe(addr, staticMux)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
