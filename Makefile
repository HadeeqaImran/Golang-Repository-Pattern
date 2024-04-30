.DEFAULT_GOAL := all

.PHONY: all generate run

all: generate run 

generate:
	swag init

run:
	go run main.go

env:
	export GOPATH=$(HOME); \
	export GOBIN=$(GOPATH)/bin; \
	export PATH=$(PATH):$(GOBIN);