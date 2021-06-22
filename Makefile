LANG = en_US.UTF-8
SHELL = /bin/bash
.SHELLFLAGS = -eu -o pipefail -c # run '/bin/bash ... -c /bin/cmd'
.DEFAULT_GOAL = build

.PHONY: build
build: ## Build ewallet
	go build

.PHONY: install
install: ## Install ewallet binary to local system
	go install

.PHONY: clean
clean: ## Remove all files created by this Makefile
	rm -rf \
		ewallet

.PHONY: help
help: ## Show Help
	@grep -E '^[a-zA-Z_-]+%?:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-20s %s\n", $$1, $$2}'|sort

