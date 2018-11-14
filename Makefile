makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

bootstrap:
	brew install go
	brew install dep
	go get github.com/UnnoTed/fileb0x
	go get github.com/githubnemo/CompileDaemon

init:
	dep ensure

purge:
	rm -rf vendor/
	rm -rf system/ab0x.go

build:
	go generate
	go build
