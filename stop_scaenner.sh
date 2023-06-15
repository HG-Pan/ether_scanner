#!/bin/bash

CONTAINER_NAME="ether_scanner"


# 检查容器是否存在
if docker container inspect "$CONTAINER_NAME" >/dev/null 2>&1; then
  echo "停止容器 $CONTAINER_NAME"
  docker container stop "$CONTAINER_NAME" >/dev/null
  echo "移除容器 $CONTAINER_NAME"
  docker container rm "$CONTAINER_NAME" >/dev/null
else
  echo "容器 $CONTAINER_NAME 不存在，无需停止"
fi

