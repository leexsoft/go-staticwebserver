/*
	纯静态站点的WEB服务器
	脚本，样式表，图片等静态资源使用http.Handle注册
	自定义路由根据传入的路径自动匹配，返回对应的静态网页文件，无需注册多个handler
	静态文件路径支持无后缀名的路径，有后缀名的路径仅支持.htm,.html，每个目录下的默认首页为index.htm
*/
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var staticMux = &StaticMux{sve: myServer}

type StaticMux struct {
	sve *server
}

//注册静态资源文件的路由
func (this *StaticMux) RegisterStaticResource() {
	for _, folder := range this.sve.GetStaticFolders() {
		http.Handle(folder, http.FileServer(http.Dir(this.sve.App.Root)))
	}
}

// 实现Handler接口
func (this *StaticMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rawUrl := r.URL.String()
	fmt.Printf("request '%s'", rawUrl)
	fmt.Println()

	if len(r.URL.RawQuery) > 0 {
		fmt.Printf("request '%s' is a dymanic requst", rawUrl)
		fmt.Println()
		http.NotFound(w, r)
		return
	}

	if r.Method == "GET" {
		if this.isStaticFile(rawUrl) {
			//使用http.Handle注册的静态资源都保存在DefaultServeMux的map内
			//需要调用DefaultServeMux来查找handler进行请求
			http.DefaultServeMux.ServeHTTP(w, r)
		} else {
			this.staticHandler(w, r)
		}
	}
}

// 判断是否是静态资源文件
func (this *StaticMux) isStaticFile(rawUrl string) bool {
	for _, ext := range this.sve.GetStaticFileExtensionds() {
		if strings.HasSuffix(rawUrl, ext) {
			return true
		}
	}

	return false
}

// 自动匹配路径的handler
func (this *StaticMux) staticHandler(w http.ResponseWriter, r *http.Request) {
	pathHtml := this.buildParseFilePath(r.URL.Path)
	if len(pathHtml) > 0 {
		fmt.Printf("ParseFile path is '%s'", (this.sve.App.Root + pathHtml))
		fmt.Println()

		t, err := template.ParseFiles(this.sve.App.Root + pathHtml)
		if err != nil {
			fmt.Println("static page parse error:", err.Error())
			http.NotFound(w, r)
		} else {
			w.Header().Set("content-type", "text/html")
			t.Execute(w, nil)
		}
	} else {
		http.NotFound(w, r)
	}
}

// 自动解析请求路径，生成html文件的模版路径
func (this *StaticMux) buildParseFilePath(rawPath string) string {
	//结尾为"/"，添加默认页面index
	pathHtml := ""
	if strings.HasSuffix(rawPath, "/") {
		pathHtml = rawPath + "index"
	} else {
		pathHtml = rawPath
	}

	//取最后一个"/"的部分
	lstIdx := strings.LastIndexAny(pathHtml, "/")
	lstSegment := pathHtml[lstIdx:]
	if strings.IndexAny(lstSegment, ".") < 0 {
		return pathHtml + ".htm"
	} else if strings.HasSuffix(lstSegment, ".html") || strings.HasSuffix(lstSegment, ".htm") || strings.HasSuffix(lstSegment, ".shtml") {
		return pathHtml
	} else {
		return ""
	}
}
