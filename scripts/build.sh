#!/usr/bin/env sh

set -ex

VERSION=$(cat version)

echo $VERSION

LDFLAGS="-X main.Version=${VERSION}"
UNAME_S=$(uname -s)
GOOS=linux
if [ "$UNAME_S" == "Darwin" ]; then
	GOOS=darwin
fi

go fmt ./...
GOOS=${GOOS} GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o feature-show-server ./main.go