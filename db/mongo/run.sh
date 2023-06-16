#!/bin/bash

# 获取当前脚本所在目录的路径
loc="$(cd "$(dirname "$BASH_SOURCE")" && pwd)"

NETWORK_NAME="location_network"

# 检查网络是否存在
if ! docker network inspect "$NETWORK_NAME" >/dev/null 2>&1; then
  # 创建网络
  docker network create "$NETWORK_NAME"
  echo "创建了 $NETWORK_NAME 网络"
else
  echo "$NETWORK_NAME 网络已经存在"
fi

# 创建日志目录（如果不存在）
logs_dir="$loc/logs"
data_dir="$loc/data"
mkdir -p "$logs_dir"
mkdir -p "$data_dir"

# 运行 MongoDB 容器并连接到指定网络
docker run -d -p 27017:27017 --name mongodb-container \
  --network="$NETWORK_NAME" \
  -e MONGO_INITDB_ROOT_USERNAME=pan \
  -e MONGO_INITDB_ROOT_PASSWORD=pan \
  -v "$logs_dir:/var/log/mongodb" \
  -v "$data_dir:/data/db" \
  mongodb:1.0
