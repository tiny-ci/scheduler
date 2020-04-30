.DEFAULT_GOAL := run

bundle:
	GOOS=linux go build

build:
	go build

run: build
	./core
