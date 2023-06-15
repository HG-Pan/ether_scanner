#!/bin/bash

docker run -d -p 27017:27017 --name mongodb-container \
  -e MONGO_INITDB_ROOT_USERNAME=pan \
  -e MONGO_INITDB_ROOT_PASSWORD=pan \
  mongodb

