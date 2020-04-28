# go-staticwebserver

==================

本服务器是一个非常简单的WEB静态资源服务端，可用于前端站点的宿主，*推荐使用docker*

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
