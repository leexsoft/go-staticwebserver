package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

const (
	TypeDirectory = "directory" // 目录
	TypeRedirect  = "redirect"  // 重定向
)

type Config struct {
	Root  string       `json:"root"` // 根目录
	Port  int64        `json:"port"` // 端口
	Paths []ConfigItem `json:"path"` // 路径
}

type ConfigItem struct {
	Type        string `json:"type"` // 类型
	Source      string `json:"src"`  // 源地址
	Destination string `json:"dst"`  // 目标地址
}

func loadAppConfig() (conf Config) {
	// 加载配置文件
	data, err := ioutil.ReadFile("app.json")
	if err != nil {
		log.Fatalln("app.json读取失败", err)
	}
	// json反序列化
	if err = json.Unmarshal(data, &conf); err != nil {
		log.Fatalln("app.json格式错误", err)
	}
	// 初始判断
	if conf.Root == "" {
		conf.Root = "/"
	}
	return
}

func main() {
	conf := loadAppConfig()

	// 遍历路由配置
	for _, path := range conf.Paths {
		switch path.Type {
		case TypeDirectory:
			http.Handle(path.Source, http.FileServer(http.Dir(filepath.Join(conf.Root, path.Destination))))
		case TypeRedirect:
			http.HandleFunc(path.Source, func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, path.Destination, http.StatusFound)
			})
		}
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil); err != nil {
		log.Fatalln("启动HTTP服务器失败", err)
	}
}
