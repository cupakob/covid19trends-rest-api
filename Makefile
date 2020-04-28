PROJECT_NAME  = covid19trends-rest-api
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
#GOPATH   = /home/bcs/.go
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)

export GO111MODULE=on

GO      = go
GODOC   = godoc
GOFMT   = gofmt
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: fmt | $(info $(M) building executable…)  ## Build program binary
	$Q $(GO) build -o $(CURDIR)/bin/$(PROJECT_NAME) main/covid19trends-rest-api.go

$(BASE): ; $(info $(M) setting GOPATH…)

# Tools

$(BIN):
	@mkdir -p $@
$(BIN)/%: $(BIN) | $(BASE) ; $(info $(M) building $(REPOSITORY)…)
	go get $(REPOSITORY)

GOLINT = $(BIN)/golint
$(BIN)/golint: REPOSITORY=golang.org/x/lint/golint

# Tests
.PHONY: check test tests
check test tests: fmt lint | $(BASE) ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q cd $(CURDIR) && $(GO) test -timeout $(TIMEOUT)s $(ARGS) ./...

test-coverage: fmt lint | $(BASE) ; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q go test -cover ./...

.PHONY: lint
lint: $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q cd $(CURDIR) && $(GOLINT) ./...

#	-set_exit_status $(PKGS)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
#	@cd $(BASE) && $(GOFMT) -l -w ./...
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... ); do \
		gofmt -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf bin

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
