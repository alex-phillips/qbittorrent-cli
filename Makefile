VERSION=$(shell git describe --tags)
LDFLAGS=-ldflags "-X main.Version=${VERSION} -s -w"
#.PHONY: build
build:
	GOOS=windows go build -ldflags="-s -w" -o dist/qbtcli.exe *.go
	GOOS=linux go build -ldflags="-s -w" -o dist/qbtcli-linux *.go
	GOOS=darwin go build -ldflags="-s -w" -o dist/qbtcli-darwin *.go
