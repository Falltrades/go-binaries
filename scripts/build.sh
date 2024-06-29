#!/bin/bash

set -e

containerName=GO_CONTAINER
imageOS=${1}
networkModule=${2}

docker run --rm -i -d --name ${containerName} \
  --net=host \
  --mount type=bind,source="$(pwd)"/go-binaries/,destination=/tmp/go-binaries/ \
  ${imageOS} sh
cp ../${networkModule}/${networkModule}.go go-binaries/
docker exec -it ${containerName} sh -c "cd /tmp/go-binaries && go mod init curl && go mod tidy && go build ."
docker stop ${containerName}
