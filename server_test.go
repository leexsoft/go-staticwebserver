package main

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	//读取配置文件
	addr := myServer.App.IP + myServer.App.Port
	fmt.Println(addr)

	//静态资源路由注册
	staticMux.RegisterStaticResource()

	//err := http.ListenAndServe(addr, staticMux)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
