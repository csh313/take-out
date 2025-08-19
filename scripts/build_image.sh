#!/bin/bash
set -eux

declare -A params=(
    ["DockerfileName"]=""
    ["DockerfileDebugName"]=""
    ["Debug"]=""
    ["BuildOnLocalDocker"]=""
    ["EndOfArgs"]=""
)
proj_root=$(dirname $(dirname $(realpath $0)))

img_url="cr.io.plus:80/io-agent/io-agent:${gitlabMergeRequestLastCommit-test}"
function set_default_args() {
    params["DockerfileName"]="Dockerfile.release"
    params["DockerfileDebugName"]="Dockerfile.debug"
    params["Debug"]="0"
    params["BuildOnLocalDocker"]="1"
    params["EndOfArgs"]="EndOfArgs"
}
function check_empty_args() {
    for key in "${!params[@]}"; do
        if [[ -z "${params[$key]}" ]]; then
            echo "Error: Parameter '$key' must be provided and non-empty."
            exit 1
        fi
    done
}

function print_all_args() {
    set +x
    for key in "${!params[@]}"; do
        echo "$key=${params[$key]}"
    done
    set -x
}
docker_file_dir=${proj_root}/deployments

# 远端构建镜像
function build_on_remote_buildkitd(){
    dockerfileName=${params["DockerfileName"]}
    debug=${params["Debug"]}
    if [ "$debug" -eq "1" ]; then
        dockerfileName=${params["DockerfileDebugName"]}
    fi
    cd ${proj_root}
    buildctl \
    --addr  ${BUILDKITD_ADDR} \
    build \
    --progress plain \
    --export-cache type=registry,ref=${CACHE_IMG_URL} \
    --import-cache type=registry,ref=${CACHE_IMG_URL} \
    --frontend=dockerfile.v0 \
    --local context=${proj_root} \
    --local dockerfile=${docker_file_dir} \
    --opt filename=${dockerfileName} \
    --output type=image,push=true,name=${img_url}
}

function build_on_local_docker(){
    dockerfileName=${params["DockerfileName"]}
    debug=${params["Debug"]}
    if [ "$debug" -eq "1" ]; then
        dockerfileName=${params["DockerfileDebugName"]}
    fi

    cd ${proj_root}
    
    docker build \
        -f ${docker_file_dir}/${dockerfileName} \
        -t ${img_url} \
        ${proj_root}

    docker push ${img_url}
}

function build_image(){
    buildOnLocalDocker=${params["BuildOnLocalDocker"]}
    if [ "$buildOnLocalDocker" -eq "1" ]; then
        build_on_local_docker
        return
    fi
    build_on_remote_buildkitd
}




