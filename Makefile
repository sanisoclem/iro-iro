
# Go parameters
BIN = ./bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

FUNCTIONS = $(shell find functions/ -type f -name 'main.go')
FUNCTION_BINARIES =  $(patsubst functions/%/main.go, bin/function-%,  $(FUNCTIONS))


all: deps clean test build

# Build shit (phew makefiles are awesome)
build: $(FUNCTION_BINARIES)
$(FUNCTION_BINARIES) : $(BIN)
$(BIN):
	mkdir $(BIN)
bin/function-% : functions/%/main.go
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $@ $<

.PHONY:
test:
	$(GOTEST) -v ./...

.PHONY:
clean:
	$(GOCLEAN)

.PHONY:
deps:
	$(GOGET) -u github.com/aws/aws-lambda-go/...

