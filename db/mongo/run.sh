#!/bin/bash
# 创建一个自定义的Docker网络
docker network create location_network

docker run -d -p 27017:27017 --name mongodb-container \
  --network=location_network \
  -e MONGO_INITDB_ROOT_USERNAME=pan \
  -e MONGO_INITDB_ROOT_PASSWORD=pan \
  mongodb

