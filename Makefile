OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
MAKE_COLOR=\033[33;01m%-20s\033[0m

all: deps tools lint test ## install deps, lint and test

deps: ## install dependencies
	@echo "$(OK_COLOR)--> Download go.mod dependencies$(NO_COLOR)"
	go mod download
	go mod vendor

tools: ## install dev tools, linters, code generators, etc..
	@echo "$(OK_COLOR)--> Installing tools from tools/tools.go$(NO_COLOR)"
	@export GOBIN=$$PWD/tools/bin; export PATH=$$GOBIN:$$PATH; cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

lint: ## run linters
	@echo "$(OK_COLOR)--> Running linters$(NO_COLOR)"
	tools/bin/golangci-lint run

test: test-unit test-e2e ## run all tests

test-unit: ## run unit tests
	@echo "$(OK_COLOR)--> Running unit tests$(NO_COLOR)"
	go test --race --count=1 ./...

test-e2e: ## run e2e tests
	@echo "$(OK_COLOR)--> Running E2E tests$(NO_COLOR)"
	go test --tags "e2e" --race --count=1 ./tests/e2e/...

fmt: ## format go files
	@echo "$(OK_COLOR)--> Formatting go files$(NO_COLOR)"
	go fmt ./...

clean: ## remove tools
	@echo "$(OK_COLOR)--> Clean up$(NO_COLOR)"
	rm -rf $(PWD)/tools/bin

help: ## show this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  $(MAKE_COLOR) %s\n", $$1, $$2 } /^##@/ { printf "\n$(MAKE_COLOR)\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: all deps tools lint test test-unit test-e2e fmt clean help