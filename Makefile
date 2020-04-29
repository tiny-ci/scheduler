.DEFAULT_GOAL := build

build:
	go build

bundle:
	GOOS=linux go build
