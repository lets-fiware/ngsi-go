curdir = $(shell pwd)
container = ngsi-cmd-build
env = -e GITHUB=$(NGSIGO_GITHUB) -e USER_ID=$(shell id -u) -e GROUP_ID=$(shell id -g) -e NGSIGO_TARGET
vol = -v $(curdir)/../build:/build

build:
	docker build -t $(container) .
run:
	docker run -it --rm $(env) $(vol) $(container)
