GO ?= go
GOFMT ?= gofmt "-s"
GOGET ?= go get "-u"
GOINSTALL ?= go install
GOBUILD ?= go build
GOTEST ?= go test

PACKAGES ?= $(shell $(GO) list ./...)
GOFILES := $(shell find . -name "*.go")
LDFLAGS ?= -ldflags="-s -w"

BINARY_NAME ?= rest-go
PKG_NAME ?= github.com/ks6088ts/$(BINARY_NAME)
CMD_DIR ?= cmd/$(BINARY_NAME)
COBRA_CONFIG ?= .cobra.yml
OUTPUT_DIR ?= outputs
CMD ?= hello
PARENT_CMD ?= rootCmd
PACKAGE_DIR ?= ./$(CMD_DIR)

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

.PHONY: fmt
fmt: ## format codes
	$(GOFMT) -w $(GOFILES)

.PHONY: lint
lint: ## run golint
	@hash golint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GOGET) golang.org/x/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

.PHONY: vet
vet: ## run vet
	for PKG in $(PACKAGES); do go vet $$PKG || exit 1; done;

.PHONY: tidy
tidy: ## tidy go modules
	go mod tidy

.PHONY: build
build: ## build applications
	$(GOBUILD) $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME) $(PACKAGE_DIR)

.PHONY: test
test: ## test go modules
	$(GOTEST) -cover -v ./...

.PHONY: ci
ci: fmt vet lint build test ## run ci

.PHONY: clean
clean: ## clean outputs
	rm -rf $(OUTPUT_DIR)

# ---
# cobra
# ---

.PHONY: install-cobra
install-cobra:
	@hash cobra > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GOGET) github.com/spf13/cobra/cobra; \
	fi

.PHONY: cobra-init
cobra-init: ## initialize cobra cli
	mkdir -p $(CMD_DIR) && \
	cd $(CMD_DIR) && \
	cobra init \
		--pkg-name $(PKG_NAME)/$(CMD_DIR) \
		--config ../../$(COBRA_CONFIG)

.PHONY: cobra-add
cobra-add: ## add cobra command
	cd $(CMD_DIR) && \
	cobra add $(CMD) \
		--config ../../$(COBRA_CONFIG) \
		--parent $(PARENT_CMD)
