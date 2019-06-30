GIT_REF := $(shell git describe --always --tag)
VERSION ?= commit-$(GIT_REF)

LINT_TOOLS=\
	golang.org/x/lint/golint \
	go get -u github.com/client9/misspell/cmd/misspell \
	github.com/kisielk/errcheck

.PHONY: all
all: build lint test

.PHONY: bootstrap-lint-tools
bootstrap-lint-tools:
	@for tool in $(LINT_TOOLS) ; do \
		echo "Installing/Updating $$tool" ; \
		go get -u $$tool; \
	done

.PHONY: get-dep
get-dep:
	@GO111MODULE=off go get github.com/golang/dep/cmd/dep

.PHONY: dep
dep: get-dep
	@GO111MODULE=off dep ensure -v

.PHONY: lint
lint:
	@go fmt
	@golint
	@misspell
	@errcheck

.PHONY: build
build: dep
	CGO_ENABLED=0 go build -o bin/yakiire \
        -ldflags "-X main.version=$(VERSION)"

.PHONY: test
test:
	@docker-compose up -d firestore
	@sleep 2
	@FIRESTORE_EMULATOR_HOST=localhost:8080 go test -v ./...
	@docker-compose stop firestore
