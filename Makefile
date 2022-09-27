.PHONY: build

GIT_COMMIT := $(shell git rev-list -1 HEAD)

##go build -ldflags="-X 'main.Version=v2.1.0'"

build:
	go build -ldflags "-X main.gitCommit=$(GIT_COMMIT)" .
run:
	./testVersioning.exe

