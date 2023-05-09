DEFAULT_GOAL := help
.PHONY:

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54323 dbname=chat user=chat-user password=chat-password sslmode=disable"

# HELP =================================================================================================================
# This will output the help for each task with comment
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

install-go-deps: ## Install deps
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

goose-install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

generate_grpc: ## Generate grpc api
	mkdir -p "pkg/chat_v1"
	protoc --proto_path=api/chat_v1 \
 		--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=./bin/protoc-gen-go \
		--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
    	--plugin=protoc-gen-go-grpc=./bin/protoc-gen-go-grpc \
    	api/chat_v1/service.proto

generate_grpc_win: ## Generate grpc api using windows protoc binaries
	mkdir -p "pkg/chat_v1"
	protoc --proto_path=api/chat_v1 \
 		--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=./bin/protoc-gen-go.exe \
		--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
    	--plugin=protoc-gen-go-grpc=./bin/protoc-gen-go-grpc.exe \
    	api/chat_v1/service.proto

local-migration-status: ## Migration status
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up: ## Migration up
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down: ## Migration down
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v
