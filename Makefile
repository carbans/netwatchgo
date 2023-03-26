.PHONY: build clean tool lint help

swag:
	swag init

run:
	go run main.go

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specific go tool"
	@echo "make clean: remove object files and cached files"