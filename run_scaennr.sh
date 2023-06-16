#!/bin/bash

# 获取当前脚本所在目录的路径
loc="$(cd "$(dirname "$BASH_SOURCE")" && pwd)"

# 创建日志目录（如果不存在）
logs_dir="$loc/logs"
mkdir -p "$logs_dir"

# 定义容器名称
container_name="ether_scanner"

# 检查镜像是否存在
if docker image inspect ether_scanner:1.0 >/dev/null 2>&1; then
    # 镜像已存在，直接运行容器
    docker run --network=location_network -d -v "$logs_dir":/app/logs --name "$container_name" ether_scanner:1.0
else
    # 镜像不存在，先构建镜像再运行容器
    docker build -t ether_scanner:1.0 "$loc" && docker run --network=location_network -d -v "$logs_dir":/app/logs --name "$container_name" ether_scanner:1.0
fi

docker ps

echo "ether_scanner Startup complete..."



