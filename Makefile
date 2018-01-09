SHELL:=/bin/bash

APPNAME =`basename ${PWD}`

build:
	@go build
run: build
	@./$(APPNAME)
