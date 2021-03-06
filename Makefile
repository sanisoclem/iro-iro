
# Go parameters
BIN = ./bin
GOCMD=@go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

FUNCTIONS = $(shell find lambda/ -type f -name 'main.go')
FUNCTION_BINARIES =  $(patsubst lambda/%/main.go, bin/function_%,  $(FUNCTIONS))

all: deps clean test build

# Build shit (phew makefiles are awesome)
build: $(FUNCTION_BINARIES)
$(FUNCTION_BINARIES) : $(BIN)
$(BIN):
	mkdir $(BIN)
bin/function_% : lambda/%/main.go
	export GOOS=linux
	export GOARCH=amd64
	$(GOBUILD) -x -o $@ $<

.PHONY:
test:
	$(GOTEST) -v ./...

.PHONY:
clean:
	$(GOCLEAN)
	@rm -rf ./bin

.PHONY:
deps:
	$(GOGET) -v -u github.com/aws/aws-lambda-go/...
	$(GOGET) -v -u github.com/aws/aws-sdk-go/aws
	$(GOGET) -v -u github.com/aws/aws-sdk-go/aws/session
	$(GOGET) -v -u github.com/aws/aws-sdk-go/service/dynamodb/...
	$(GOGET) -v -u github.com/rs/xid
	$(GOGET) -v -u github.com/gusaul/go-dynamock