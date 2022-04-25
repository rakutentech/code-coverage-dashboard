#!/bin/bash

source ~/.bashrc

kill -9 `cat ./pids/*.pid`
./main $port > /dev/null 2>&1 &
