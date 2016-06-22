COLOR=\033[32;01m
NO_COLOR=\033[0m

NAME := stan-challenge
TAGS := $(shell git describe --always --dirty --tags)
LDFLAGS := "-X main.VERSION=$(TAGS)"

.PHONY: setup vet lint fmt build test

all: test build

setup:
	@echo "$(COLOR)==> Setting up deps...$(NO_COLOR)"
	@go get -u "github.com/tools/godep"
	@go get -u "github.com/golang/lint/golint"
	@go get -u "github.com/kisielk/errcheck"
	@echo

errcheck:
	@echo "$(COLOR)==> errcheck$(NO_COLOR)"
	@echo

vet:
	@echo "$(COLOR)==> go vet$(NO_COLOR)"
	@go vet ./...
	@echo

lint:
	@echo "$(COLOR)==> golint$(NO_COLOR)"
	@golint .
	@echo

fmt:
	@echo "$(COLOR)==> go fmt$(NO_COLOR)"
	@go fmt ./...
	@echo

test: fmt vet lint errcheck
	@echo "$(COLOR)==> go test$(NO_COLOR)"
	@godep go test ./... -cover
	@echo

build:
	@echo "$(COLOR)==> go build$(NO_COLOR)"
	@mkdir -p bin/
	@godep go build -ldflags=$(LDFLAGS) -o bin/$(NAME)
	@echo

clean:
	@echo "$(COLOR)==> Cleaning up bin/...$(NO_COLOR)"
	@rm -rfv bin/*
	@echo
