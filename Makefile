LANG = en_US.UTF-8
SHELL = /bin/bash
.SHELLFLAGS = -eu -o pipefail -c # run '/bin/bash ... -c /bin/cmd'
.DEFAULT_GOAL = build

GOIMPORTS = $(GOPATH)/bin/goimports
STATICCHECK = $(GOPATH)/bin/staticcheck
GOLANGCI-LINT = $(GOPATH)/bin/golangci-lint

.PHONY: build
build: ## Build ewallet
	go build

.PHONY: install
install: ## Install ewallet binary to local system
	go install

$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports@latest

$(STATICCHECK):
	go install honnef.co/go/tools/cmd/staticcheck@latest

$(GOLANGCI-LINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1

fmt:
	gofmt -w -s .

goimports: fmt $(GOIMPORTS)
	goimports -w .

staticcheck: $(STATICCHECK)
	staticcheck -go 1.17 ./...

golangci-lint: $(GOLANGCI-LINT)
	golangci-lint run ./...

.PHONY: clean
clean: ## Remove all files created by this Makefile
	rm -rf \
		ewallet

.PHONY: help
help: ## Show Help
	@grep -E '^[a-zA-Z_-]+%?:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-20s %s\n", $$1, $$2}'|sort

