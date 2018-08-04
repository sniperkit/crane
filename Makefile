APP_PREFIX := "crane"
APP_SUFFIX := ""
APP_NAME := "$(APP_PREFIX)$(SEPARATOR)$(APP_DIRNAME)$(APP_SUFFIX)"

PWD = $(shell pwd)
SEPARATOR := "-"
INSTALL_DIR := "/usr/local/bin/"
# DIST_DIRS := find * -type d -exec
DIST_DIR := "../../../dist"
BIN_DIR := "../../../bin"

# determine platform
ifeq (Darwin, $(findstring Darwin, $(shell uname -a)))
  PLATFORM := darwin
else
  PLATFORM := linux
endif

PLATFORM_VERSION ?= $(shell uname -r)
PLATFORM_ARCH ?= $(shell uname -m)
PLATFORM_INFO ?= $(shell uname -a)
PLATFORM_OS ?= $(shell uname -s)

APP_ROOT := $(shell pwd)
APP_DIRNAME := $(shell basename `pwd`)
APP_PKG_URI ?= $(shell pwd | sed "s\#$(GOPATH)/src/\#\#g")
APP_PKG_URI_ARRAY ?= $(shell pwd | sed "s\#$(GOPATH)/src/\#\#g" | tr "/" "\n")

GO_EXECUTABLE ?= $(shell which go)
GO_VERSION ?= $(shell $(GO_EXECUTABLE) version)

GOX_EXECUTABLE ?= $(shell which gox)
GOX_VERSION ?= "master"

GLIDE_EXECUTABLE ?= $(shell which glide)
GLIDE_VERSION ?= $(shell $(GLIDE_EXECUTABLE) --version)

GODEPS_EXECUTABLE ?= $(shell which dep)
GODEPS_VERSION ?= $(shell $(GODEPS_EXECUTABLE) version | tr -s ' ')

GIT_EXECUTABLE ?= $(shell which git)
GIT_VERSION ?= $(shell $(GIT_EXECUTABLE) version)

SEPARATOR := "-"
INSTALL_DIR := "/usr/local/bin/"
# DIST_DIRS := find * -type d -exec
DIST_DIR := "$(CURDIR)/dist"
BIN_DIR := "$(CURDIR)/bin"


APP_ROOT := $(shell pwd)
APP_DIRNAME := $(shell basename `pwd`)
APP_PKG_URI ?= $(shell pwd | sed "s\#$(GOPATH)/src/\#\#g")
APP_PKG_URI_ARRAY ?= $(shell pwd | sed "s\#$(GOPATH)/src/\#\#g" | tr "/" "\n")
APP_PKG_DOMAIN ?= "$(word 1, $(APP_PKG_URI_ARRAY))"
APP_PKG_OWNER ?= "$(word 2, $(APP_PKG_URI_ARRAY))"
APP_PKG_NAME ?= "$(word 3, $(APP_PKG_URI_ARRAY))"
APP_PKG_URI_ROOT ?= "$(APP_PKG_DOMAIN)/$(APP_PKG_OWNER)/$(APP_PKG_NAME)"
APP_PKG_LOCAL_PATH ?= "$(GOPATH)/src/$(APP_PKG_URI_ROOT)"
APP_SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
APP_PREFIX := $(shell basename $(APP_PKG_LOCAL_PATH))
APP_SUFFIX := ""
APP_NAME := "$(APP_PREFIX)$(SEPARATOR)$(APP_DIRNAME)$(APP_SUFFIX)"

VERSION ?= $(shell git describe --tags)
VERSION_INCODE = $(shell perl -ne '/^var version.*"([^"]+)".*$$/ && print "v$$1\n"' main.go)
VERSION_INCHANGELOG = $(shell perl -ne '/^\# Release (\d+(\.\d+)+) / && print "$$1\n"' CHANGELOG.md | head -n1)

VCS_GIT_REMOTE_URL = $(shell git config --get remote.origin.url)
VCS_GIT_VERSION ?= $(VERSION)

# docker commands:
# - docker build -t sniperkit/crane:go1.10.3-alpine3.7-dev --no-cache .
# - docker build -t sniperkit/crane:go1.10.3-alpine3.7-console -f Dockerfile.console .
# - docker build -t sniperkit/crane:go1.10.3-alpine3.7-prod --target=runner .
# - docker build -t sniperkit/crane:go1.10.3-alpine3.7-dist --target=xcross .
# - docker build -t sniperkit/crane:go1.10.3-debian-wheezy-dist -f Dockerfile.xcross .

.PHONY: print-% fmt default
print-%: ; @echo $*=$($*)

.PHONY: all build install build-dist
all: build install build-dist ## build and install crane package locally, then, build all dist versions

test: ## run all unit tests in this package
	@(go list ./... | grep -v "vendor/" | xargs -n1 go test -v -cover)

fmt: ## format source code with gofmt
	@(gofmt -w crane)

.PHONY: docker docker-pull docker-push docker-xcross docker-console

docker: docker-pull
	@docker build -t sniperkit/crane:go1.10.3-alpine3.7-dev .

docker-all: docker-console docker-xcross

docker-runner:
	@docker build -t sniperkit/crane:go1.10.3-alpine3.7-prod --target=runner .

docker-runner:
	@docker build -t sniperkit/crane:go1.10.3-alpine3.7-prod --target=runner .

docker-console:
	@docker build -t sniperkit/crane:go1.10.3-alpine3.7-console -f Dockerfile.console .

docker-xcross: docker
	@docker build -t sniperkit/crane:go1.10.3-debian-wheezy-dist --target=xcross .
	@docker run -ti --rm -e CGO_ENABLED=0 \
	-v $(CURDIR):/gopath/src/github.com/sniperkit/crane \
	-w /gopath/src/github.com/sniperkit/crane \
	sniperkit/crane:go1.10.3-debian-wheezy-dist \
	gox \
	-osarch="darwin/amd64 darwin/386 linux/amd64 linux/386 windows/amd64 windows/386" \
	-output "dist/{{.Dir}}_{{.OS}}_{{.Arch}}"

docker-xcross-standalone:
	@docker build -t sniperkit/crane:go1.10.3-debian-wheezy-standalone -f Dockerfile.xcross .

docker-push:
	@docker push sniperkit/crane:go1.10.3-debian-wheezy-dist
	@docker push sniperkit/crane:go1.10.3-alpine3.7-dev


# $(shell date -v-1d +%Y-%m-%d)
build: build-$(PLATFORM) ## build local executable of the default crane cli version
default: build-$(PLATFORM)

.PHONY: run run-linux run-darwin run-darwin-pro run-windows-pro run-windows-pro
run: run-$(PLATFORM)

.PHONY: build-linux build-darwin build-darwin-pro build-windows build-windows-pro
build-dist: build-linux build-darwin build-darwin-pro build-windows build-windows-pro ## build crane for all platforms

.PHONY: install-linux install-darwin
install: install-$(PLATFORM) ## install crane for your local platform

run-linux: ## run crane for linux (64bits)
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go run ./cmd/crane/*.go

build-linux: ## build crane for linux (64bits)
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/crane -v github.com/sniperkit/crane/cmd/crane
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(DIST_DIR)/crane_linux_amd64 -v github.com/sniperkit/crane/cmd/crane

install-linux:
	@go install github.com/sniperkit/crane/cmd/crane
	@crane version

build-darwin: ## build crane for MacOSX (64bits)
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/crane -v github.com/sniperkit/crane/cmd/crane
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(DIST_DIR)/crane_linux_amd64 -v github.com/sniperkit/crane/cmd/crane

run-darwin: ## run crane for MacOSX (64bits)
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go run -race ./cmd/crane/*.go

install-darwin: ## install crane for MacOSX (64bits)
	@go install github.com/sniperkit/crane/cmd/crane
	@crane version

build-darwin-pro: ## build crane pro for MacOSX (64bits)
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -tags pro -o $(DIST_DIR)/crane_darwin_amd64_pro -v github.com/sniperkit/crane/cmd/crane

run-darwin-pro: ## build crane pro for MacOSX (64bits)
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go run ./cmd/crane/*.go

build-windows: ## build crane for Windows (64bits)
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o $(DIST_DIR)/crane_windows_amd64.exe -v github.com/sniperkit/crane/cmd/crane

run-windows: ## run crane for Windows (64bits)
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go run ./cmd/crane/*.go

build-windows-pro: ## build crane pro for Windows (64bits)
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags pro -o $(DIST_DIR)/crane_windows_amd64_pro.exe -v github.com/sniperkit/crane/cmd/crane

run-windows-pro: ## run crane pro for Windows (64bits)
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go run ./cmd/crane/*.go

.PHONY: help
help: ## display available makefile targets for this project
	@echo "\033[36mMAKEFILE TARGETS:\033[0m"
	@echo "- PLATFORM: $(PLATFORM)"
	@echo "- PWD: $(PWD)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "- \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)