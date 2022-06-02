#!/bin/bash

source ~/.bashrc

cd server
kill -9 `cat ./pids/*.pid`
./main 3006 > /dev/null 2>&1 &
