#VERSION=$(shell git describe --tags)
LDFLAGS=-ldflags "-X main.Version=${VERSION} -s" # -w" (https://golang.org/cmd/link/)
#.PHONY: build
build:
	GO111MODULE=on GOOS=windows go build $(LDFLAGS) -o dist/qbtcli.exe *.go
	GO111MODULE=on GOOS=linux go build $(LDFLAGS) -o dist/qbtcli-linux *.go
	GO111MODULE=on GOOS=darwin go build $(LDFLAGS) -o dist/qbtcli-darwin *.go
