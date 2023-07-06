#!/bin/bash
RUN_NAME=openiot-gateway
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
CGO_ENABLED=0 go build \
     -ldflags " \
     -X 'main.BuildVersion=${build_version}' \
     -X 'main.BuildGoVersion=${go_version}' \
     -X 'main.BuildTime=${build_time}' \
     -X 'main.BuildCommit=${build_commit}' \
     -o output/bin/${RUN_NAME}
     " main.go