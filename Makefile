# Makefile for NOORCHAIN Core
#
# This file provides a few simple helper commands to build and inspect
# the NOORCHAIN node binary (`noord`).
#
# All commands are meant to be run in a Go-enabled environment
# (for you or CI). You do NOT need to run them locally if you don't want to.

BINARY_NAME = noord

.PHONY: all build clean

all: build

## build: build the NOORCHAIN node binary
build:
	go build -o bin/$(BINARY_NAME) ./cmd/noord

## clean: remove build artifacts
clean:
	rm -rf bin
