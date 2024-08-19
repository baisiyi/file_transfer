FROM golang:1.18

WORKDIR /file_transfer

COPY . .

RUN go mod tidy  \
    && GOOS=linux GOARCH=arm64 go build -o file_transfer .

# 设置容器启动时的默认命令
CMD ["./file_transfer"]