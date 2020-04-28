# go-staticwebserver

==================

本服务器是一个非常简单的WEB静态资源服务端，主要作为前端工程的宿主服务端，*推荐使用docker实现*

## 代码结构

* `/main`：主程序
* `/app.json`：配置文件

## 配置文件

* 配置文件必须与可执行文件在同一目录下
* 配置文件名必须为 app.json
* 根目录的默认值为`/`,root节可不设值
* 端口的默认值为`80`,port节必须设值

```json
{
    "root": "/Users/leexsoft/Documents/Code/leexsoft/github/go/staticwebserver/html",
    "port": 80,
    "path": [
        {"type": "directory", "src": "/", "dst": ""},
        {"type": "directory", "src": "/css", "dst": "/css"},
        {"type": "directory", "src": "/js", "dst": "/js"},
        {"type": "redirect", "src": "/login", "dst": "/"}
    ]
}
```

## docker操作

打包命令

```shell
docker build -t leexsoft/staticweb:latest .
```

在Docker中，包含如下资源：

* `/staticweb`：主程序
* `/app.json`：配置文件，从具体的前端工程中传入
* `/html/`：推荐的静态资源根目录，可在配置文件中配置，也是从具体的前端工程中传入
