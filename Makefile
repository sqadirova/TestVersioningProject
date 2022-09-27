.PHONY: build

#GIT_COMMIT := $(shell git rev-list -1 HEAD)

build:
	go build -ldflags="-X 'main.Version=v2.1.0'"
run:
	./testVersioning.exe

