# replace BINARY_NAME with your service name
BINARY_NAME=elastic-search-example
GOCMD=go

.PHONY: build
build:
	@echo ------Go build from make file---------
	mkdir -p bin/$(BINARY_NAME)
	GO111MODULE=on $(GOCMD) build -o bin/$(BINARY_NAME)/$(BINARY_NAME) ./cmd/