#!/usr/bin/env bash

ADDRESS=$1
if [[ ${ADDRESS} -eq "" ]]; then
    ADDRESS=127.0.0.1:6060
fi
# 用 -u 分析cpu使用情况，请确认go-torch在可执行目录下
go-torch -u http://${ADDRESS}

