#!/bin/bash
set -eux

wait_for_mysql() {
    echo "Waiting for MySQL..."
    local count=0
    local timeout=60
    while ! mysql -h"mysql" -P"3306" -uroot -p"123456" -e "SELECT 1" &>/dev/null; do
        count=$((count + 1))
        if [ $count -gt $timeout ]; then
            echo "MySQL not ready after $timeout seconds, aborting"
            exit 1
        fi
        sleep 1
    done
    if ! mysql -h"mysql" -P"3306" -uroot -p"123456" -e "USE hmshop" &>/dev/null; then
        echo "Creating database hmshop..."
        mysql -h"mysql" -P"3306" -uroot -p"123456" -e "CREATE DATABASE IF NOT EXISTS hmshop"
    fi
    echo "MySQL is ready with database hmshop"
}

wait_for_redis() {
    echo "Waiting for Redis..."
    local count=0
    local timeout=30
    
    while ! redis-cli -h "redis" ping | grep -q "PONG"; do
        count=$((count + 1))
        if [ $count -gt $timeout ]; then
            echo "Redis not ready after $timeout seconds, aborting"
            exit 1
        fi
        sleep 1
    done
    echo "Redis is ready"
}

wait_for_mysql
wait_for_redis

# 执行后续命令
exec "$@"