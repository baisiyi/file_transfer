FROM golang:1.18

# 设置 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /file_transfer

COPY . .

RUN go mod tidy  \
    && GOOS=linux GOARCH=arm64 go build -o file_transfer .

# 设置容器启动时的默认命令
CMD ["./file_transfer"]