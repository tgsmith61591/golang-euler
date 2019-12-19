SHELL := /bin/bash

ROOT_MAKEFILE_PATH := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

.PHONY: run
run:
	go run cmd/main.go

.PHONY: help
help:
	@echo "Available commands:"
	@echo run
	@echo
