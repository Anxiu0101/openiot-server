include ../common.mk

SERVICE_NAME = user

.PHONY: server
server:
	sh build.sh
	cd output && sh bootstrap.sh

.PHONY: gen
gen:
	kitex \
	-gen-path ../../kitex_gen \
	-service "$(SERVICE_NAME)" \
	-module "$(MODULE)" \
	-type thrift \
	../../idl/$(SERVICE_NAME).thrift
	go mod tidy

.PHONY: test
test:
	cd test && go test -v


#.PHONY: generate
#generate:
#	kitex \
#    -service "openiot-user" \
#    -module "github.com/OpenIoT-Hub/openiot-server" \
#    -type protobuf \
#    -I ./../../idl/ \
#    -I ./../../idl/user/ user.proto
#
#.PHONY: run_client
#run_client:
#	go run ./client/main.go
#
#.PHONY: run_server
#run_server:
#	go run main.go handler.go

# Service: user
# Path: internal/user
# Optional: -use "github.com/OpenIoT-Hub/openiot-server/kitex_gen"
# kitex -service "openiot-user" -module "github.com/OpenIoT-Hub/openiot-server" \
# -use "github.com/OpenIoT-Hub/openiot-server/kitex_gen" \
# -type protobuf -I ./../../idl/ device.proto