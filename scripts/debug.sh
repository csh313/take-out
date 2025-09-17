#!/bin/bash
set -e

echo "Starting debug mode..."


# 构建应用
echo "Building application..."
go build -gcflags="all=-N -l" -o /app/server main.go

echo "Starting Delve debugger on port 1234..."
# 启动 delve 调试器，使用 DAP 模式
dlv --listen=:1234 --headless=true --api-version=2 --accept-multiclient --continue exec /app/server