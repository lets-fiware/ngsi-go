# MIT License
# 
# Copyright (c) 2020-2021 Kazuhito Suda
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
VERSION := $(shell sed -n -r '/var version/s/.*"(.*)".*/\1/p' cmd/ngsi/main.go)
REVISION := $(shell git rev-parse HEAD)
LDFLAGS := -X main.revision=$(REVISION) -w -s
NAMEVER := $(NAME)-v$(VERSION)

build = rm -f build/$(3) && CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) go build -ldflags "$(LDFLAGS)" -o build/$(3) cmd/$(NAME)/main.go
buildx = script/release_buildx.sh build/$(1).tar.gz
tar = cd build && tar -cvzf $(NAMEVER)-$(1)-$(2).tar.gz $(3) && rm $(3)
sum = cd build && sha256sum *.tar.gz > $(NAMEVER)-sha256sum.txt

export GO111MODULE=on

.PHONY: devel-deps
devel-deps:
	go get \
	github.com/tcnksm/ghr \
	golang.org/x/tools/cmd/cover

bin/%: cmd/%/main.go $(SOURCES)
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $@ $<

.PHONY: unit-test
unit-test:
	script/unit-test.sh

.PHONY: coverage
coverage:
	script/coverage.sh

.PHONY: golangci-lint
golangci-lint:
	script/golangci-lint.sh

.PHONY: e2e-test
e2e-test:
	script/e2e-test.sh

.PHONY: lint-dockerfile
lint-dockerfile:
	script/lint-dockerfile.sh

.PHONY: lint-docs
lint-docs:
	script/lint-docs.sh

.PHONY: all-tests
all-tests:
	script/all-tests.sh

.PHONY: wc
wc:
	wc $(SOURCES)

.PHONY: build
build: bin/ngsi

.PHONY: release
release:
	@rm -fr build/
	@script/release_buildx.sh
	$(call sum)

.PHONY: release-github
release-github:
	ghr -c $(REVISION) v$(VERSION) build/

.PHONY: build-all
build-all: linux darwin 

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
darwin: build/darwin_amd64.tar.gz build/darwin_arm64.tar.gz

build/darwin_amd64.tar.gz: $(SOURCES)
	$(call build,darwin,amd64,$(NAME))
	$(call tar,darwin,amd64,$(NAME))

build/darwin_arm64.tar.gz: $(SOURCES)
	$(call build,darwin,arm64,$(NAME))
	$(call tar,darwin,arm64,$(NAME))

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

.PHONY: darwin_arm64
darwin_arm64:
	$(call buildx,$@)

.PHONY: version
version:
	@echo $(VERSION)
