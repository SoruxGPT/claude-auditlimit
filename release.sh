#!/bin/bash

set -e

gf build main.go -a amd64 -s linux -p ./temp
gf docker main.go -p -t epicmo/claude-auditlimit:latest
now=$(date +"%Y%m%d%H%M%S")
# 以当前时间为版本号
docker tag epicmo/claude-auditlimit:latest epicmo/claude-auditlimit:$now
docker push epicmo/claude-auditlimit:$now
echo "release success" $now
