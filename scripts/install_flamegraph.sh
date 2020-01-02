#!/usr/bin/env bash

# install go-torch
go get github.com/uber/go-torch

# install flamegraph
cd /opt/ && git clone https://github.com/brendangregg/FlameGraph.git
ln -sf /opt/FlameGraph/flamegraph.pl /usr/local/bin/flamegraph

# install graphviz 还需要安装一个graphviz用来画内存图
apt-get install graphviz -y


