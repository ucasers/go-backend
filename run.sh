#!/bin/bash

# 拉取最新镜像
docker pull ucas789/icampus:latest

# 停止并删除旧容器（如果存在）
docker stop go_server
docker rm go_server

# 后台运行新容器
docker run -d -p 10085:10085 --name go_server ucas789/icampus:latest