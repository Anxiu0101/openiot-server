.PHONY: generate
generate:
	kitex \
	-service "openiot-notice" \
	-module "github.com/OpenIoT-Hub/openiot-server" \
	-type protobuf \
	-I ./../../idl/ notice.proto

.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: run_server
run_server:
	go run main.go handler.go

# Service: notice
# Path: internal/notice
# Optional: -use "github.com/OpenIoT-Hub/openiot-server/kitex_gen"
# kitex -service "openiot-user" -module "github.com/OpenIoT-Hub/openiot-server" \
# -use "github.com/OpenIoT-Hub/openiot-server/kitex_gen" \
# -type protobuf -I ./../../idl/ notice.proto
