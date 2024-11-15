#!/bin/bash

set -x

SOURCE_DIR=`pwd`/pkg/$1
BUILD_DIR=${BUILD_DIR:-${SOURCE_DIR}/build}
BUILD_TYPE=${BUILD_TYPE:-release}
INSTALL_DIR=${INSTALL_DIR:-`pwd`/libs}
CXX=${CXX:-g++}

mkdir -p ${BUILD_DIR} \
    && cd ${BUILD_DIR} \
    && cmake \
        -DCMAKE_BUILD_TYPE=${BUILD_TYPE} \
        -DCMAKE_INSTALL_PREFIX=${INSTALL_DIR} \
        -DCMAKE_EXPORT_COMPILE_COMMANDS=ON \
        $SOURCE_DIR \
    && make -j8 && make install \
    && rm -rf ${BUILD_DIR}
