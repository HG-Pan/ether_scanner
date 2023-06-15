# 基于官方的Go镜像作为基础
FROM golang:1.19.4

# 设置容器内日志文件的路径
ENV LOG_FILE_PATH=/app/logs/mylog.log

# 将当前目录的所有内容复制到容器的 /app 目录
COPY . /app

# 设置工作目录为 /app
WORKDIR /app

# 构建Go程序
RUN go build -o main

# 修改启动命令，将日志输出重定向到指定的日志文件
CMD ["./main", "2>&1", "|", "tee", "-a", "$LOG_FILE_PATH"]
