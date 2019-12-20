#!/usr/bin/env bash

ADDRESS=$1
if [[ ${ADDRESS} -eq "" ]]; then
    ADDRESS=127.0.0.1:8888
fi

# 用 -alloc_space 分析内存的临时分配情况
go-torch -alloc_space http://${ADDRESS}/debug/pprof/heap --colors=mem


