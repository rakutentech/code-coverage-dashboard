#!/bin/bash

source ~/.bashrc

cd client
ps aux|grep 3009|grep -v grep|awk '{print $2}'|xargs kill -9
pnpm start > /dev/null 2>&1 &
