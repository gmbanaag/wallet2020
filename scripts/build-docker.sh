#!/bin/sh

set -e

if [ $# -ne 1 ]
then
    echo "Usage: $0 [SERVICE NAME]"
    exit
fi

TAG="latest"
SERVICE_NAME=${1}

echo "building $SERVICE_NAME"

echo "using tag $TAG"

make linux

docker build -t ${SERVICE_NAME}:${TAG} . -f ./build/package/Dockerfile
