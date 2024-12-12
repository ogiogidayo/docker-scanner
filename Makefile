SHELL=/bin/bash
BIN_DIR := $(shell pwd)/bin
GOIMPORTS := $(abspath $(BIN_DIR)/goimports)
GOLANGCI_LINT := $(abspath $(BIN_DIR)/golangci-lint)

$(BIN_DIR):
	mkdir -p ./bin

goimports: $(BIN_DIR) $(GOIMPORTS)
$(GOIMPORTS):
	cd ./tools/goimports && go build -o ../../bin/goimports golang.org/x/tools/cmd/goimports

golangci-lint: $(BIN_DIR) $(GOLANGCI_LINT)
$(GOLANGCI_LINT):
	cd ./tools/golangci-lint && go build -o ../../bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: fmt
fmt: $(GOIMPORTS)
	find . -print | grep --regex '.*\.go$$' | xargs $(GOIMPORTS) -w -local "github.com/ogiogidayo/docker-scanner"

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run -v ./...

.PHONY: test
test:
	go test -v ./...