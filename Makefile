LOCAL_BIN:=$(CURDIR)/bin

### Linter functional
install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2

lint:
	./bin/golangci-lint run ./... --config .golangci.pipeline.yaml

### Protobugg functional
install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	mkdir -p pkg/swagger
	make generate-chat_server-api
	$(LOCAL_BIN)/statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

generate-chat_server-api:
	mkdir -p pkg/chat_server_v1
	protoc --proto_path api/chat_server_v1 --proto_path vendor.protogen \
	--go_out=pkg/chat_server_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_server_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--grpc-gateway_out=pkg/chat_server_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--validate_out lang=go:pkg/chat_server_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/chat_server_v1/chat_server.proto

### Docker service deploy
build-linux:
	GOOS=linux GOARCH=amd64 go build -o service_chat cmd/server/main.go

copy-to-server:
	scp service_linux_chat_server root@31.129.49.166:

docker-build-and-push-registry:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/test/chat-server:v0.0.1 .
	docker login -u token -p CRgAAAAAkMI2zCW2BiycXtSp2ufvWNw3pimuCJow cr.selcloud.ru/test/chat-server:v0.0.1
	docker push cr.selcloud.ru/test/chat-server:v0.0.1

# docker pull cr.selcloud.ru/test/chat-server:v0.0.1
# docker run -p 50552:50552 cr.selcloud.ru/test/chat-server:v0.0.1

### Goose functional
include local.env

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DB_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

local-migration-status:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

deploy-all-local:
	docker-compose up --build -d

install-minimock:
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@latest

.PHONY: test
test:
	go clean -testcache
	go test ./... -v

.PHONY: test-coverage
test-coverage:
	go clean -testcache
	-go test ./... -v -coverprofile=coverage.tmp.out -covermode count -coverpkg=github.com/markgenuine/chat-server/internal/service/...,github.com/markgenuine/chat-server/internal/api/... -count 5
	grep -v "mocks/" coverage.tmp.out  > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=./coverage.out | grep "total";

vendor-proto:
	@if [ ! -d vendor.protogen/validate ]; then \
		mkdir -p vendor.protogen/validate &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
		rm -rf vendor.protogen/protoc-gen-validate ;\
	fi
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
		mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
		git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
		mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
		rm -rf vendor.protogen/openapiv2 ;\
	fi