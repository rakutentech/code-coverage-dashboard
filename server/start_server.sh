#!/bin/bash

source ~/.bashrc

port=$1
kill -9 `cat ./pids/*.pid`
./main $port > /dev/null 2>&1 &
