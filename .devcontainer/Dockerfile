FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

SHELL ["/bin/bash", "-c"]

# 2. 安装必要工具
RUN apt-get update && apt-get install -y \
    git \
    make \
    curl \
    default-mysql-client \
    redis-tools \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# 3. 设置Go环境
ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=on \
    GOPATH=/go \
    PATH=$PATH:/go/bin:/app

WORKDIR /app
COPY . .

# 4. 添加数据库检查脚本
COPY scripts/wait-for-db.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/wait-for-db.sh

RUN chown -R root:root /app /go

USER root
CMD ["bash"]