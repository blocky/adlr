GOFLAGS=GOFLAGS=-mod=mod
GO=$(GOFLAGS) go1.16.3
GOTEST=$(GO) test -count=1
GOMOD=$(GO) mod
GOTIDY=$(GOMOD) tidy

INTEGRATION=integration
TIMEOUT=5m

MOCK=mockery
MOCKS=internal/mocks

default: test

mock: # autogenerate mocks for interface testing
	@$(MOCK) --all --output=./$(MOCKS)

test: test-unit test-integration

test-integration:
	@$(GOTEST) -timeout=$(TIMEOUT) ./$(INTEGRATION)/... \
	&& $(GOTIDY)

test-unit:
	@$(GOTEST) -short ./... && $(GOTIDY)