#!/usr/bin/env bash

# 画出内存分配图
go tool pprof -alloc_space -cum -svg http://${ADDRESS}

