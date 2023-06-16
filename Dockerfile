# 使用带有Alpine Linux的Golang镜像作为构建环境
FROM golang:1.19.4-alpine as builder

RUN apt-get update && apt-get install -y procps vim

# 设置时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

# 设定工作目录
WORKDIR /app

# 将你的代码复制到Docker环境中
COPY . .

# 下载依赖并编译Go程序
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用带有Alpine Linux的Golang镜像作为基础镜像
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

# 从构建环境中复制二进制文件到当前环境
COPY --from=builder /app/main /app/main

RUN mkdir -p /app/logs

# 设定环境变量
ENV LOG_FILE_PATH=/app/logs/app.log

# 启动服务
ENTRYPOINT ["/app/main"]
