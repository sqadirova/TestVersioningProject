GOOS=linux
GOARCH=386

.PHONY: build

GIT_COMMIT := $(shell git rev-list -1 HEAD)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.gitCommit=$(GIT_COMMIT)" .


