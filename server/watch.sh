#!/bin/bash

reflex -R "assets|circleci|.git|logs" -r '\.go$' -s -- sh -c 'go run main.go 3000'