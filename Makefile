NAME := git-bucket-to-lab
BINARY_NAME = $(NAME)
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) -v

clean:
	go clean
	rm -f $(BINARY_NAME)

.PHONY: build clean
