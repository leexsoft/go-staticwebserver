# 打包镜像
FROM golang:1.13.6-alpine3.11 as builder
WORKDIR /src
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-w -s" -o dist/staticweb

# 运行时镜像
FROM xiaochengtech/alpine-timezone:latest
# 拷贝可执行文件
COPY --from=builder /src/dist/staticweb /staticweb
RUN chmod +x /staticweb
# 设置运行环境
WORKDIR /
EXPOSE 80
ENTRYPOINT [ "/staticweb" ]
