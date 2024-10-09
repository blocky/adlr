GO=go
GOTEST=$(GO) test -count=1
GOMOD=$(GO) mod
GOTIDY=$(GOMOD) tidy

MOCK=mockery
MOCKS=internal/mocks

default: test

mock: # autogenerate mocks for interface testing
	@$(MOCK) --all --output=./$(MOCKS)

test:
	@$(GOTEST) ./... && $(GOTIDY)

