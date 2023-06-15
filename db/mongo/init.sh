#!/bin/bash

# Check if an existing Docker image with the same name exists
if [[ "$(docker images -q mongodb 2> /dev/null)" != "" ]]; then
    echo "Existing image with the same name found, replacing..."
    # Remove the existing Docker image with the same name
    docker rmi mongodb
fi

# Build the Docker image
docker build -t mongodb .
