FROM golang:1.18

WORKDIR /file_transfer

RUN go mod tidy  \
    && make \

COPY . .

RUN chmod +x file_transfer

# 设置容器启动时的默认命令
CMD ["./file_transfer"]