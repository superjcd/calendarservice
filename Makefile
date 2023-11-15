.DEFAULT_GOAL := all 

# =============================================================================
# Globals:
ROOT_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
OUTPUT_DIR  := $(ROOT_DIR)/_output
TEST_DIRS = $(ROOT_DIR)/service/implement
PLATFORMS ?=   linux_amd64 windows_amd64
VERSION = 1.0
GO_MODULE = $(word 2, $(shell type go.mod))  
COMMAND = $(lastword $(subst /, , $(GO_MODULE_NAME)))

# =============================================================================

.PHONY: all 
all: prepare tidy test build publish

.PHONY: prepare
prepare:
	-@mkdir _output
	@echo "====>outpur directory prepared sucessfully"

.PHONY: genproto
genproto:
	@protoc --go_out=./ --go_opt=module=$(GO_MODULE) --go-grpc_out=./ --go-grpc_opt=module=$(GO_MODULE) ./proto/*.proto

.PHONY: tidy 
tidy:
	@go mod tidy

.PHONY: test
test:
	@go test  --coverprofile=$(OUTPUT_DIR)/coverage.out  $(TEST_DIRS)
	@go tool cover --html=$(OUTPUT_DIR)/coverage.out -o=$(OUTPUT_DIR)/coverage.html


.PHONY: build
build:  $(foreach P,${PLATFORMS}, $(addprefix build., $(P)))

.PHONY: build.%
build.%:
	$(eval OS:= $(word 1,$(subst _, ,$*)))
	$(eval ARCH := $(word 2,$(subst _, ,$*)))  
	$(if $(findstring windows, $(OS)), $(eval EXE_SUFFIX:=.exe), $(eval EXE_SUFFIX:=''))
	@go env -w CGO_ENABLED=0  GOOS=$(OS) GOARCH=$(ARCH)
	@echo "====>Build binary for ${COMMAND}, with OS: $(OS), ARCH:$(ARCH)"
	@go build -o $(OUTPUT_DIR)/$(COMMAND)_$(OS)_$(ARCH)$(EXE_SUFFIX)  $(ROOT_DIR)/cmd/


.PHONY: clean
clean:
	$(if $(findstring Windows, $(OS)), $(shell rmdir /Q /S _output), $(shell rm -rf _output))
	@echo "====>output directory  is removed sucessfully"


.PHONY: publish
publish: image.build image.push

.PHONY: image.build 
image.build:
	@echo "====>Build docker images"
	@docker build -t superjcd/${COMMAND} .


.PHONY: image.push
image.push:
	@docker push superjcd/${COMMAND}:latest