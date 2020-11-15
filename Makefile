# MIT License
# 
# Copyright (c) 2020 Kazuhito Suda
# 
# This file is part of NGSI Go
# 
# https://github.com/lets-fiware/ngsi-go
# 
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
# 
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
# 
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

NAME := ngsi
SOURCES := cmd/ngsi/main.go internal/ngsicmd/*.go internal/ngsilib/*.go
VERSION := $(shell gobump show -r cmd/$(NAME))
REVISION := $(shell git rev-parse HEAD)
LDFLAGS := -X main.revision=$(REVISION) -w -s
NAMEVER := $(NAME)-v$(VERSION)

build = rm -f build/$(3) && GOOS=$(1) GOARCH=$(2) go build -ldflags "$(LDFLAGS)" -o build/$(3) cmd/$(NAME)/main.go
buildx = script/release_buildx.sh build/$(1).tar.gz
tar = cd build && tar -cvzf $(NAMEVER)-$(1)-$(2).tar.gz $(3) && rm $(3) && sha256sum $(NAMEVER)-$(1)-$(2).tar.gz > $(NAMEVER)-$(1)-$(2)-sha256sum.txt

export GO111MODULE=on

.PHONY: devel-deps
devel-deps:
	go get \
	github.com/x-motemen/gobump/cmd/gobump \
	github.com/tcnksm/ghr

bin/%: cmd/%/main.go $(SOURCES)
	go build -ldflags "$(LDFLAGS)" -o $@ $<

.PHONY: test
test:
	cd internal;go test -coverprofile=profile.out -covermode=atomic ./...

.PHONY: cover
cover:
	cd internal;go test -coverprofile=cover.out -covermode=atomic ./...;go tool cover -html=cover.out -o cover.html

.PHONY: wc
wc:
	wc $(SOURCES)

.PHONY: build
build: bin/ngsi

.PHONY: all clean linux drrwin

.PHONY: release
release:
	@rm -fr build/
	@script/release_buildx.sh

.PHONY: release_github
release_github:
	ghr -c $(REVISION) v$(VERSION) build/

.PHONY: all
all: linux darwin 

.PHONY: clean
clean:
	@rm -rf build/

.PHONY: linux
linux: build/linux_amd64.tar.gz build/linux_arm64.tar.gz build/linux_arm.tar.gz

build/linux_amd64.tar.gz: $(SOURCES)
	$(call build,linux,amd64,$(NAME))
	$(call tar,linux,amd64,$(NAME))

build/linux_arm64.tar.gz: $(SOURCES)
	$(call build,linux,arm64,$(NAME))
	$(call tar,linux,arm64,$(NAME))

build/linux_arm.tar.gz: $(SOURCES)
	$(call build,linux,arm,$(NAME))
	$(call tar,linux,arm,$(NAME))

.PHONY: darwin
darwin: build/darwin_amd64.tar.gz

build/darwin_amd64.tar.gz: $(SOURCES)
	$(call build,darwin,amd64,$(NAME))
	$(call tar,darwin,amd64,$(NAME))

.PHONY: linux_amd64
linux_amd64:
	$(call buildx,$@)

.PHONY: linux_arm64
linux_arm64:
	$(call buildx,$@)

.PHONY: linux_arm
linux_arm:
	$(call buildx,$@)

.PHONY: darwin_amd64
darwin_amd64:
	$(call buildx,$@)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: bump_major
bump_major:
	@gobump major -w cmd/$(NAME)

.PHONY: bump_minor
bump_minor:
	@gobump minor -w cmd/$(NAME)

.PHONY: bump_patch
bump_patch:
	@gobump patch -w cmd/$(NAME)
