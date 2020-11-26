GOCMD    = go
GOBUILD  = $(GOCMD) build
GOTEST   = $(GOCMD) test

all: build

.PHONY: build run

build:
	$(GOBUILD)

run: build
	./sqli-demo-server
