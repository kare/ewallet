LANG = en_US.UTF-8
SHELL = /bin/bash
.SHELLFLAGS = -e -u -o pipefail -c
.DEFAULT_GOAL = build

GOIMPORTS = $(GOPATH)/bin/goimports
STATICCHECK = $(GOPATH)/bin/staticcheck

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

.PHONY: fmt
fmt:
	gofmt -w -s .

.PHONY: goimports
goimports: $(GOIMPORTS)
	goimports -w .

.PHONY: staticcheck
staticcheck: $(STATICCHECK)
	staticcheck ./...

.PHONY: clean
clean: ## Remove all files created by this Makefile
	go clean -i

.PHONY: help
help: ## Show Help
	@grep -E '^[a-zA-Z_-]+%?:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-20s %s\n", $$1, $$2}'|sort

