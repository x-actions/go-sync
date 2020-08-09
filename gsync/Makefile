.PHONY: all test clean build build-linux build-mac build-windows

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=gsync
FILE="./VERSION"
VERSION=$(shell cat $(FILE))
DIR_LINUX=$(BINARY_NAME)_$(VERSION)_linux
DIR_MAC=$(BINARY_NAME)_$(VERSION)_mac
DIR_WIN=$(BINARY_NAME)_$(VERSION)_win

all: clean test build build-linux build-mac build-windows
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -rf bin/*
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v

linux:
	if [ ! -d bin/$(DIR_LINUX) ]; then mkdir -p bin/$(DIR_LINUX); fi
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(DIR_LINUX)/$(BINARY_NAME) -v
	cp README.md ../LICENSE bin/$(DIR_LINUX)/;cd bin; tar -zcvf $(DIR_LINUX).tar.gz $(DIR_LINUX)
mac:
	if [ ! -d bin/$(DIR_MAC) ]; then mkdir -p bin/$(DIR_MAC); fi
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(DIR_MAC)/$(BINARY_NAME) -v
	cd bin; tar -zcvf $(DIR_MAC).tar.gz $(DIR_MAC)
windows:
	if [ ! -d bin/$(DIR_WIN) ]; then mkdir -p bin/$(DIR_WIN); fi
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/$(DIR_WIN)/$(BINARY_NAME) -v
	cd bin; tar -zcvf $(DIR_WIN).tar.gz $(DIR_WIN)
