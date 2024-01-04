BINDIR      := $(CURDIR)/bin
INSTALL_PATH ?= /usr/local/bin
DIST_DIRS   := find * -type d -exec
BINNAME     ?= esctl
TARGETS = amd64-darwin arm64-darwin amd64-linux arm64-linux amd64-windows

GOBIN         = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         = $(shell go env GOPATH)/bin
endif
GOX           = $(GOBIN)/gox
GOIMPORTS     = $(GOBIN)/goimports
ARCH          = $(shell uname -p)

# go option
PKG         := ./...
TAGS        :=
TESTS       := .
TESTFLAGS   :=
LDFLAGS     := -w -s
GOFLAGS     :=
CGO_ENABLED ?= 0

GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)

ifdef VERSION
	BINARY_VERSION = $(VERSION)
endif
BINARY_VERSION ?= ${GIT_TAG}

.PHONY: deps
deps:
	go mod tidy
	go mod download

.PHONY: build
build:
	GO111MODULE=on CGO_ENABLED=$(CGO_ENABLED) go build $(GOFLAGS) -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(BINNAME) main.go

.PHONY: build-cross
build-cross: LDFLAGS += -extldflags "-static"
build-cross: $(foreach target,$(TARGETS),build-$(target))

define build-target
build-$(1):
	GOFLAGS="-trimpath" GOARCH=$(word 1,$(subst -, ,$(1))) GOOS=$(word 2,$(subst -, ,$(1))) GO111MODULE=on CGO_ENABLED=0 go build -p 3 -o _dist/$(1)/$(BINNAME) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' main.go
endef

$(foreach target,$(TARGETS),$(eval $(call build-target,$(target))))

_dist:
	mkdir -p _dist

.PHONY: dist
dist:
	( \
		cd _dist && \
		$(DIST_DIRS) cp ../README.md {} \; && \
		$(DIST_DIRS) tar -zcf $(BINNAME)-${BINARY_VERSION}-{}.tar.gz {} \; \
	)

.PHONY: clean
clean:
	@ggrm -rf ./_dist


.PHONY: linter
linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$GOPATH/bin

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: run
run:
	go run main.go

.PHONY: dk-build
dk-build:
	docker build -t esctl:latest .

.PHONY: mockgen
mockgen:
	@go get go.uber.org/mock/mockgen@latest
	mockgen -source ./internal/file/editor.go -destination ./internal/file/editor_mock.go -package file
	mockgen -source ./internal/client/elastic.go -destination ./internal/client/elastic_mock.go -package client

.PHONY: test
test:
	go test ./... -covermode=count -count 1
