SHELL := bash
NAME := controller
BIN := bin
DIST := dist

# Determine the OS and set the executable name accordingly
ifeq ($(OS), Windows_NT)
	EXECUTABLE := $(NAME).exe
	UNAME := Windows
else
	EXECUTABLE := $(NAME)
	UNAME := $(shell uname -s)
endif

# Build settings
GOBUILD ?= CGO_ENABLED=0 go build
PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f -not -path ./.devenv/\* -not -path ./.direnv/\*)
GENERATE ?= $(PACKAGES)
TAGS ?= netgo

# Output and version settings
ifndef OUTPUT
	ifeq ($(GITHUB_REF_TYPE), tag)
		OUTPUT ?= $(subst v,,$(GITHUB_REF_NAME))
	else
		OUTPUT ?= testing
	endif
endif

ifndef VERSION
	ifeq ($(GITHUB_REF_TYPE), tag)
		VERSION ?= $(subst v,,$(GITHUB_REF_NAME))
	else
		VERSION ?= $(shell git rev-parse --short HEAD)
	endif
endif

ifndef DATE
	DATE := $(shell date -u '+%Y%m%d')
endif

ifndef SHA
	SHA := $(shell git rev-parse --short HEAD)
endif

# Compiler flags
LDFLAGS += -s -w -extldflags "-static" -X "$(IMPORT)/pkg/version.String=$(VERSION)" -X "$(IMPORT)/pkg/version.Revision=$(SHA)" -X "$(IMPORT)/pkg/version.Date=$(DATE)"
GCFLAGS += all=-N -l

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Examines Go source
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

.PHONY: vet
vet:
	go vet $(PACKAGES)

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Linters
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

.PHONY: staticcheck
staticcheck: $(STATICCHECK)
	$(STATICCHECK) -tags '$(TAGS)' $(PACKAGES)

.PHONY: lint
lint: $(REVIVE)
	for PKG in $(PACKAGES); do $(REVIVE) -config revive.toml -set_exit_status $$PKG || exit 1; done;

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Simplifies code generation
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

.PHONY: generate
generate:
	go generate $(GENERATE)

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Testing
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

.PHONY: test
test:
	go test -coverprofile coverage.out $(PACKAGES)

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Build
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

.PHONY: build
build: $(BIN)/$(EXECUTABLE)

$(BIN)/$(EXECUTABLE): $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(BIN)/$(EXECUTABLE)-debug: $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -gcflags '$(GCFLAGS)' -o $@ ./

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Release
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

# Release targets
.PHONY: release
release: $(DIST) release-linux release-darwin release-windows release-checksum

$(DIST):
	mkdir -p $(DIST)

# Linux release targets
.PHONY: release-linux
release-linux: $(DIST) \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-386 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-amd64 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-5 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-6 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-7 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm64 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-mips \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-mips64 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-mipsle \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-mips64le

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-386:
	GOOS=linux GOARCH=386 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@  ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-5:
	GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-6:
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-7:
	GOOS=linux GOARCH=arm GOARM=7 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-mips:
	GOOS=linux GOARCH=mips $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-mips64:
	GOOS=linux GOARCH=mips64 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-mipsle:
	GOOS=linux GOARCH=mipsle $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-mips64le:
	GOOS=linux GOARCH=mips64le $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

# Darwin release targets
.PHONY: release-darwin
release-darwin: $(DIST) \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-darwin-amd64 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-darwin-arm64

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

# Windows release targets
.PHONY: release-windows
release-windows: $(DIST) \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-windows-4.0-386.exe \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-windows-4.0-amd64.exe

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-windows-4.0-386.exe:
	GOOS=windows GOARCH=386 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-windows-4.0-amd64.exe:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./

.PHONY: release-checksum
release-checksum:
	cd $(DIST); $(foreach file,$(wildcard $(DIST)/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)