GOFLAGS=GOFLAGS=-mod=mod
GO=$(GOFLAGS) go
GOTEST=$(GO) test -count=1
GOMOD=$(GO) mod
GOTIDY=$(GOMOD) tidy
GOLIST=$(GO) list

INTEGRATION=integration
TIMEOUT=5m

MOCK=mockery
MOCKS=internal/mocks

BUILDLIST=buildlist.json

default: test

clean:
	@rm -rf $(BUILDLIST)

mock: # autogenerate mocks for interface testing
	@$(MOCK) --all --output=./$(MOCKS)

test: test-unit test-integration

test-integration: gen-buildlist
	@$(GOTEST) -timeout=$(TIMEOUT) ./$(INTEGRATION)/... \
	&& $(GOTIDY)

test-unit:
	@$(GOTEST) -short ./... && $(GOTIDY)

gen-buildlist:
	@$(GOLIST) -m -json all > $(BUILDLIST)
