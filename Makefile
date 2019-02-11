VERSION=$(shell git describe --tags)
LDFLAGS=-ldflags "-X main.Version=${VERSION} -s -w"
#.PHONY: build
build:
	GO111MODULE=on GOOS=windows go build -ldflags="-s -w" -o dist/qbtcli.exe *.go
	GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o dist/qbtcli-linux *.go
	GO111MODULE=on GOOS=darwin go build -ldflags="-s -w" -o dist/qbtcli-darwin *.go
