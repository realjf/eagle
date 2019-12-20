#!/usr/bin/env bash

ADDRESS=$1
if [[ ${ADDRESS} -eq "" ]]; then
    ADDRESS=127.0.0.1:8888
fi
# 用 -inuse_space 来分析程序常驻内存的占用情况
go-torch -inuse_space http://${ADDRESS}/debug/pprof/heap --colors=mem

