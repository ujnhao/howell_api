#!/bin/bash
RUN_NAME=coder.hao.howell_api
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}