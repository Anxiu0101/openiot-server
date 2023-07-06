#!/usr/bin/env bash
RUN_NAME="openiot-user"

mkdir -p output/bin
cp script/* output/
chmod +x output/bootstrap.sh

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    CGO_ENABLED=0 go build \
         -ldflags " \
         -X 'main.BuildVersion=${build_version}' \
         -X 'main.BuildGoVersion=${go_version}' \
         -X 'main.BuildTime=${build_time}' \
         -X 'main.BuildCommit=${build_commit}' \
         -o output/bin/${RUN_NAME}
         " main.go
else
    go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi

