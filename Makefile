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

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-chat_server-api

generate-chat_server-api:
	mkdir -p pkg/chat_server_v1
	protoc --proto_path api/chat_server_v1 \
	--go_out=pkg/chat_server_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_server_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
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