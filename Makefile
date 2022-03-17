BINARY_NAME := basic-app
PACKAGE_NAME := github.com/af-go/basic-app

GOCMD :=go
GOBUILD :=$(GOCMD) build
SWAGGERCMD := swagger

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GODIST=$(GOBASE)/dist

$(info argument BUILD_NUM: ${BUILD_NUM})
$(info argument BUILD_USER: ${BUILD_USER})
$(info argument GOPROXY: ${GOPROXY})

# export GOPROXY 
GO_PROXY := GOPROXY=${GOPROXY}

PKGS := $(shell ${GO_PROXY} go list ./... | grep -v vendor | grep -v testing )
GOSECFLAGS := -fmt=json -out=gosec-results.json ./...
FILES := $(shell find . -type f -name '*.go' ! -path "./vendor/*")

LDFLAGS += \
    -w -s -extldflags 'static' \
	-X '${PACKAGE_NAME}/pkg/version.Commit=$$(git rev-parse HEAD)' \
	-X '${PACKAGE_NAME}/pkg/version.GoVersion=$$(go version)' \
	-X '${PACKAGE_NAME}/pkg/version.BuildNum=${BUILD_NUM}' \
	-X '${PACKAGE_NAME}/pkg/version.BuildBy=${BUILD_USER}' \
	-X '${PACKAGE_NAME}/pkg/version.BuildAt=`date -u +%Y%m%d%H%M%S`'
BUILD_OPTIONS ?= -a -installsuffix cgo -a -tags netgo -ldflags "${LDFLAGS}"

all: build
build:
	@echo "==> building go code $(GOPROXY)"
	$(GO_PROXY) CGO_ENABLED=0 $(GOBUILD) ${BUILD_OPTIONS} -v -o ${GODIST}/${BINARY_NAME} main.go


dep:
	@echo "==> downloading go tools $(GOPROXY)"
	GOBIN=$(GOBIN) go get -u golang.org/x/tools/cmd/goimports
	GOBIN=$(GOBIN) go get -u golang.org/x/lint/golint
	GOBIN=$(GOBIN) go get github.com/securego/gosec/v2/cmd/gosec
	GOBIN=$(GOBIN) go get github.com/golang/mock/mockgen@v1.4.4

lint:
	@echo "==> lint $(GOPROXY)"
	$(GOBIN)/golint $(PKGS)

go-download:
	@echo "==> downloading deps $(GOPROXY)"
	GOBIN=$(GOBIN) go mod download
	
gosec:
	@echo "==> static files scan (gosec), see gosec-results.json for details"
	$(GOBIN)/gosec $(GOSECFLAGS)

check-fmt:
	@echo "==> check format $(GOPROXY)"
	$(GOBIN)/goimports -d $(FILES)

fmt:
	@echo "==> formating $(GOPROXY)"
	$(GOBIN)/goimports -w .

test:
	@echo "==> testing $(GOPROXY)"
	${GOCMD} test ./... -v
