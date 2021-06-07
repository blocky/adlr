GOFLAGS=GOFLAGS=-mod=mod
GO=$(GOFLAGS) go
GOTEST=$(GO) test -count=1
GOMOD=$(GO) mod
GOTIDY=$(GOMOD) tidy
GOLIST=$(GO) list
GOBUILD=$(GO) build

INTEGRATION=integration
TIMEOUT=5m

ADLR_SRC=pkg
MOCK=mockery
MOCKS=internal/mocks

BIN=bin
SCRIPTS=sh

LICENSELOCK=license.lock
BUILDLIST=buildlist.json

default: test

clean:
	@rm -rf $(BUILDLIST)
	@rm -rf $(BIN)

# mock autogeneration for interface testing
mock: mock-internal mock-pkg

mock-internal:
	@$(MOCK) --dir=./internal --all --output=./$(MOCKS)

mock-pkg:
	@$(MOCK) --dir=./$(ADLR_SRC) --all --output=./$(MOCKS)

# building
bin:
	@mkdir $(BIN)

build: build-linux-amd64

build-tmp: bin # build tmp exec to perform adlr steps
	@$(GOBUILD) -o $(BIN)/tmp .

build-linux-amd64: licenselock
	@$(SCRIPTS)/build.sh \
	adlr linux amd64 \
	. ./$(BIN) ./$(LICENSELOCK)

buildlist:
	@$(GOLIST) -m -json all > $(BUILDLIST)

licenselock: build-tmp buildlist
	@$(BIN)/tmp evaluate --buildlist=$(BUILDLIST)

# testing
test: test-unit test-integration

test-integration: buildlist
	@$(GOTEST) -timeout=$(TIMEOUT) ./$(INTEGRATION)/... \
	&& $(GOTIDY)

test-unit:
	@$(GOTEST) -short ./... && $(GOTIDY)
