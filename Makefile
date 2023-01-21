# *** params ***

APP_RELATIVE_PATH = $(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
INTERNAL_PROTO_FILES = $(shell find internal -name *.proto)
GEN_PROTO_FILES = $(shell find api -name "*.go")

APP_NAME = $(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
APP_VERSION=latest
APP_IMAGE="github.com/fxkt-tech/raiden/"$(APP_NAME):$(APP_VERSION)

BUILD_PREFIX_LINUX_AMD64 = "CGO_ENABLED=0 GOOS=linux GOARCH=amd64"
COLOR = "\e[1;36m%s\e[0m\n"

# *** fast key ***

.PHONY: genall
genall: api
	cd app/feed/service && make config
	cd app/user/service && make config
	cd app/message/service && make config

.PHONY: clean
clean: clean-api


# *** build ***

.PHONY: build
build:
	@printf $(COLOR) "Build: [$(APP_RELATIVE_PATH)]"
	@mkdir -p bin
	@go build -ldflags "-X main.Name=$(APP_NAME) -X main.Version=$(APP_VERSION)" -o ./bin/ ./...

.PHONY: build-linux-amd64-bin
build-linux-amd64-bin:
	@printf $(COLOR) "Build: [$(APP_RELATIVE_PATH)] for docker linux"
	@mkdir -p bin
	@$(BUILD_PREFIX_LINUX_AMD64) go build -ldflags "-X main.Name=$(APP_NAME) -X main.Version=$(APP_VERSION)" -o ./bin/ ./...

# *** check ***

lint:
	@printf $(COLOR) "Lint: golangci-lint"
	@golangci-lint run

.PHONY: buf-lint
buf-lint:
	@printf $(COLOR) "Lint: buf-lint"
	@buf lint

# *** code gen ***

.PHONY: wire
wire:
	@printf $(COLOR) "Codegen: [$(APP_RELATIVE_PATH)] by wire"
	@cd cmd/server && wire

.PHONY: ent
ent:
	@printf $(COLOR) "Codegen: [$(APP_RELATIVE_PATH)] by ent"
	@cd internal/data && ent generate ./ent/schema

.PHONY: api
api: clean-api
	@printf $(COLOR) "Codegen: api"
	@buf generate --template buf.gen.yaml

.PHONY: config
config:
	@printf $(COLOR) "Codegen: [$(APP_RELATIVE_PATH)] for config"
	@protoc --proto_path=. \
           --proto_path=../../../third_party \
           --go_out=paths=source_relative:. \
           $(INTERNAL_PROTO_FILES)

.PHONY: baseconfig
baseconfig:
	@printf $(COLOR) "Codegen: baseconfig"
	@cd third_party/yilan && protoc --proto_path=. \
		--go_out=paths=source_relative:. \
		./config/v1/*.proto

# *** tools ***

.PHONY: tools
tools:
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	go install entgo.io/ent/cmd/ent@v0.11.0
	cd pkg/kratos/cms/protoc-gen-go-ylres && go install

# *** clean ***

.PHONY: clean-api
clean-api:
	@printf $(COLOR) "Clean: api gen code"
	@rm -rf $(GEN_PROTO_FILES)

# *** deploy ***

PHONY: docker-build
docker-build: build-linux-amd64-bin
	@printf $(COLOR) "Deploy: docker-build [$(DOCKER_IMAGE)]"
	@cd ../../.. && docker build \
		-f deploy/build/Dockerfile \
		--build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) \
		-t $(DOCKER_IMAGE) .

PHONY: docker-push
docker-push:
	@printf $(COLOR) "Deploy: docker-push [$(DOCKER_IMAGE)]"
	@docker login --username=yilanvideo@1697015562157106 -pyilanvam2022 registry.cn-beijing.aliyuncs.com
	@cd ../../.. && docker push $(DOCKER_IMAGE)

.PHONY: test
test:
	go test -v ./... -cover
