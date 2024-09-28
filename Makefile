GOFLAGS=GOFLAGS=-mod=mod
GO=$(GOFLAGS) go
GOTEST=$(GO) test -count=1
GOMOD=$(GO) mod
GOTIDY=$(GOMOD) tidy
GORUN=$(GO) run
GOBUILD=$(GO) build

INTEGRATION=internal/integration
TIMEOUT=5m

ADLR_SRC=pkg
ADLR_MAIN=./main.go

BIN=bin
SCRIPTS=sh

VERSION=$(ADLR_MAIN)/version

GIT_TAG=$(shell git describe --tags)

default: test

clean:
	@rm -rf $(BUILDLIST)
	@rm -rf $(BIN)

# building
bin:
	@mkdir $(BIN)

build: build-linux-amd64

build-tmp: bin # build tmp exec to perform adlr steps
	@$(GOBUILD) -o $(BIN)/tmp ./$(ADLR_MAIN)

build-linux-amd64: version
	@$(SCRIPTS)/build.sh \
	adlr linux amd64 ./$(ADLR_MAIN) ./$(BIN)

lint:
	@golangci-lint run --config ./golangci.yaml

version:
	@printf $(GIT_TAG) > $(VERSION)

# testing
test: test-unit test-integration

test-integration: tidy
	@$(GOTEST) -timeout=$(TIMEOUT) ./$(INTEGRATION)/...

test-unit: tidy
	@$(GOTEST) -short ./$(ADLR_SRC)/...

tidy:
	@$(GOTIDY)
