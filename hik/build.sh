#!/bin/bash

# >>> 执行第三方包安装脚本 >>>
# ./jsoncpp_build.sh
# ./muduo_build.sh
# ./yamlcpp_build.sh
pkgs=("cpp-httplib" "yaml-cpp-master")
for pkg in "${pkgs[@]}"; do 
    ./pkg_build.sh ${pkg}
done 
# <<< 结束 <<<

set -x

SOURCE_DIR=`pwd`
PROJECT=$1
BUILD_DIR=${BUILD_DIR:-${SOURCE_DIR}/build}
BUILD_TYPE=${BUILD_TYPE:-release}
CXX=${CXX:-g++}

mkdir -p ${BUILD_DIR} \
    && cd ${BUILD_DIR} \
    && cmake \
        -DCMAKE_BUILD_TYPE=${BUILD_TYPD} \
        -DCMAKE_EXPORT_COMPILE_COMMANDS=ON \
        $SOURCE_DIR \
    && make && mv ./${PROJECT} ../${PROJECT} && rm -rf ${SOURCE_DIR}/build
