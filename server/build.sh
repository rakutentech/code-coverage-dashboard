#! /bin/sh

export https_proxy=
export http_proxy=

GOOS=linux GOARCH=amd64 go build main.go
