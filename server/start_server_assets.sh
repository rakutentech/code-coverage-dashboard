#!/bin/bash

source ~/.bashrc

caddy file-server --root ./assets --listen localhost:3008 --browse > /dev/null 2>&1 &

