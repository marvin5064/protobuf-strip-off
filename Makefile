SHELL:=/bin/bash

APPNAME =`basename ${PWD}`

default: run

build:
	@go build
run: build
	@./$(APPNAME)
