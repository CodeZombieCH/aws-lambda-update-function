.DEFAULT_GOAL := build

.PHONY: build

BUILD_DIR = build

build:

	GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/update-linux-amd64 ./cmd/update
	GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/update-win-amd64 ./cmd/update
	GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/update-macos-amd64 ./cmd/update
