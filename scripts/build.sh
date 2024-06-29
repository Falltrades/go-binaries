#!/bin/bash

set -e

tempdir=tempdir
containerName=GO_CONTAINER
imageOS=${1}
networkModule=${2}

rm -rf ${tempdir}
mkdir ${tempdir}

docker run --rm -i -d --name ${containerName} \
  --net=host \
  --mount type=bind,source="$(pwd)"/${tempdir}/,destination=/tmp/${tempdir}/ \
  ${imageOS} sh
cp ../${networkModule}/${networkModule}.go ${tempdir}
docker exec -it ${containerName} sh -c "cd /tmp/${tempdir} && go mod init ${networkModule} && go mod tidy && go build ."
docker stop ${containerName}
