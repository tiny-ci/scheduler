ARTIFACT = tiny-ci-core

.DEFAULT_GOAL := run

bundle:
	@GOOS=linux go build -o $(ARTIFACT)
	@zip $(ARTIFACT).zip $(ARTIFACT)
	@rm $(ARTIFACT)

build:
	go build

run: build
	./core
