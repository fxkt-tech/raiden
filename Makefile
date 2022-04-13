APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
API_PROTO_FILES=$(shell cd ../../../api/$(APP_RELATIVE_PATH) && find . -name *.proto)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
# APP_VERSION=$(shell head -n 1 VERSION)
APP_VERSION=latest
APP_IMAGE="fxkt.tech/raiden/"$(APP_NAME):$(APP_VERSION)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.1
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.3

.PHONY: errors
# generate errors code
errors:
	cd ../../../api/$(APP_RELATIVE_PATH) && \
	protoc --proto_path=. \
		--proto_path=../../../ \
		--proto_path=../../../third_party \
		--go_out=paths=source_relative:. \
		--go-errors_out=paths=source_relative:. \
		$(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc  --proto_path=. \
			--proto_path=../../../third_party \
			--go_out=paths=source_relative:. \
			$(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	cd ../../../api/$(APP_RELATIVE_PATH) && \
	protoc --proto_path=. \
		--proto_path=../../../ \
		--proto_path=../../../third_party \
		--go_out=paths=source_relative:. \
		--go-http_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--openapi_out==paths=source_relative:. \
		--validate_out=paths=source_relative,lang=go:. \
		$(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -o ./bin/ ./...

.PHONY: docker-build
docker-build:
	cd ../../.. && docker build \
		-f deploy/build/Dockerfile \
		--build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) \
		-t $(APP_IMAGE) .

.PHONY: docker-push
docker-push:
	cd ../../.. && docker push $(APP_IMAGE)

.PHONY: docker-compose
docker-compose:
	docker-compose -f deploy/docker-compose/docker-compose.yml -p raidenblade up

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: wire
# wire
wire:
	cd cmd/server && wire

.PHONY: pogen
# po gen
pogen:
	mkdir -p internal/data/db/query && \
	cd internal/data/db && \
	go run gen.go

.PHONY: sql
# sql
sql:
	mysql -h127.0.0.1 -uroot -pqingchuan495 -Ddb_raiden < ../../../doc/table.sql

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
