###
# Find OS and Go environment
# GO contains the Go binary
# FS contains the OS file separator
###
ifeq ($(OS),Windows_NT)
  GO := $(shell where go.exe 2> NUL)
  FS := \\
else
  GO := $(shell command -v go 2> /dev/null)
  FS := /
endif

ifeq ($(GO),)
  $(error could not find go. Is it in PATH? $(GO))
endif

GOPATH ?= $(shell $(GO) env GOPATH)
GITHUBDIR := $(GOPATH)$(FS)src$(FS)github.com
GOLANGCI_LINT_VERSION := v1.17.1
GOLANGCI_LINT_HASHSUM := f5fa647a12f658924d9f7d6b9628d505ab118e8e049e43272de6526053ebe08d

###
# Functions
###

go_get = $(if $(findstring Windows_NT,$(OS)),\
IF NOT EXIST $(GITHUBDIR)$(FS)$(1)$(FS) ( mkdir $(GITHUBDIR)$(FS)$(1) ) else (cd .) &\
IF NOT EXIST $(GITHUBDIR)$(FS)$(1)$(FS)$(2)$(FS) ( cd $(GITHUBDIR)$(FS)$(1) && git clone https://github.com/$(1)/$(2) ) else (cd .) &\
,\
mkdir -p $(GITHUBDIR)$(FS)$(1) &&\
(test ! -d $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && cd $(GITHUBDIR)$(FS)$(1) && git clone https://github.com/$(1)/$(2)) || true &&\
)\
cd $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && git fetch origin && git checkout -q $(3)

go_install = $(call go_get,$(1),$(2),$(3)) && cd $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && $(GO) install

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(shell cd $(shell dirname $(mkfile_path)); pwd)

###
# tools
###

TOOLS_DESTDIR  ?= $(GOPATH)/bin

GOLANGCI_LINT   = $(TOOLS_DESTDIR)/golangci-lint
GOIMPORTS       = $(TOOLS_DESTDIR)/goimports
RUNSIM          = $(TOOLS_DESTDIR)/runsim

all: tools

tools: tools-stamp
tools-stamp: $(GOIMPORTS) $(RUNSIM)
	touch $@

$(GOLANGCI_LINT): $(mkfile_dir)/install-golangci-lint.sh
	bash $(mkfile_dir)/install-golangci-lint.sh $(TOOLS_DESTDIR) $(GOLANGCI_LINT_VERSION) $(GOLANGCI_LINT_HASHSUM)

$(GOIMPORTS):
	go get golang.org/x/tools/cmd/goimports@v0.0.0-20190628034336-212fb13d595e

$(RUNSIM):
	go get github.com/cosmos/tools/cmd/runsim/

golangci-lint: $(GOLANGCI_LINT)

tools-clean:
	rm -f $(GOIMPORTS) $(GOLANGCI_LINT)
	rm -f tools-stamp

.PHONY: all tools tools-clean
