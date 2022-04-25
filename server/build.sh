#! /bin/sh

export https_proxy=
export http_proxy=
cp .env.prod .env

GOOS=linux GOARCH=amd64 go build main.go
