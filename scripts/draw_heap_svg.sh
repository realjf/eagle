#!/usr/bin/env bash

ADDRESS=$1
if [[ ${ADDRESS} -eq "" ]]; then
    ADDRESS=127.0.0.1:8888
fi

# 画出内存分配图
go tool pprof -alloc_space -cum -svg http://${ADDRESS}/debug/pprof/heap

