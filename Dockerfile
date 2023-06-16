# 使用官方的Golang镜像作为构建环境
FROM golang:1.19.4-alpine as builder

# 设定工作目录
WORKDIR /app

# 将你的代码复制到Docker环境中
COPY . .

# 下载依赖并编译Go程序
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Add this line to copy the SSL certificates
RUN apk --no-cache add ca-certificates
RUN update-ca-certificates

# 使用scratch作为基础镜像，构建一个没有任何额外层的镜像
FROM scratch

# 从构建环境中复制二进制文件到当前环境
COPY --from=builder /app/main /app/main

# Copy the SSL certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# 创建/app/logs文件夹存放日志文件
COPY --from=builder /app/logs /app/logs

# 设定环境变量
ENV LOG_FILE_PATH=/app/logs/app.log

# 启动服务
ENTRYPOINT ["/app/main"]
