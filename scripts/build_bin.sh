#!/bin/bash
set -eux

declare -A params=(
    ["Debug"]=""
    ["EndOfArgs"]=""
)

ldflagsG=""
gcflagsG=""

function set_default_args() {
    params["Debug"]="0"
    params["EndOfArgs"]="EndOfArgs"
}

# 解析命令行参数
function parse_args() {
    while [[ $# -gt 0 ]]; do
        case "$1" in
            --debug=*)
                params["Debug"]="${1#*=}"
                shift
                ;;
            *)
                # 未知参数
                echo "Unknown argument: $1"
                exit 1
                ;;
        esac
    done
}

# 根据模式设置编译参数
function setup_build_flags() {
    if [ "${params["Debug"]}" = "1" ]; then
        echo "Debug mode enabled"
        # 调试模式下的编译参数
        gcflagsG="all=-N -l"  # 禁用优化和内联，便于调试
        ldflagsG="-w"         # 禁用DWARF生成
    else
        echo "Run mode enabled"
        # 生产模式下的编译参数
        gcflagsG=""
        ldflagsG="-s -w"  # 缩小二进制文件大小
    fi
}

# 构建函数
function build_bin(){
    # 确保输出目录存在
    mkdir -p /app
    export CGO_ENABLED=1
    
    go mod tidy
    echo "Building binaries..."

    # 编译 server
    echo "Building server..."
    go build -v -ldflags="${ldflagsG}" -gcflags="${gcflagsG}" -o /app/server ../main.go
    ensure_success $? "build server err"
    echo "Server binary created at:"
    ls -lh /app/server

    echo "Build completed successfully!"
}

# 主函数
function main() {
    set_default_args
    parse_args "$@"
    setup_build_flags
    build_bin
}

main "$@"
