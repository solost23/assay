cmake_minimum_required(VERSION 3.0)

project(hik)
SET(HTYJ_PUBLIC_SOLUTION_ROOT_DIR "${CMAKE_CURRENT_SOURCE_DIR}")
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -I${CMAKE_SOURCE_DIR}/libs/include")

# 添加项目头文件找包目录
include_directories(${CMAKE_CURRENT_SOURCE_DIR})

# 添加项目头文件找库目录
link_directories(${CMAKE_CURRENT_SOURCE_DIR}/libs/lib)

add_executable(hik main.cpp configs/config.cpp controllers/nvr_c.cpp services/nvr_s.cpp util/error.cpp)

find_library(hcnetsdkLIBRARY hcnetsdk HINTS ${HTYJ_PUBLIC_SOLUTION_ROOT_DIR}/libs/lib)
if(hcnetsdkLIBRARY)
    message(STATUS "Found libhcnetsdk library: ${hcnetsdkLIBRARY}")
    target_link_libraries(hik ${hcnetsdkLIBRARY} ${SELF_INPUT_LIBRARIES})
else()
    message(FATAL_ERROR "Could not find libhcnetsdk library")
endif()

find_library(yamlLIBRARY yaml-cpp HINTS ${HTYJ_PUBLIC_SOLUTION_ROOT_DIR}/libs/lib)
if(yamlLIBRARY)
    message(STATUS "Found libyaml library: ${yamlLIBRARY}")
    target_link_libraries(hik ${yamlLIBRARY} ${SELF_INPUT_LIBRARIES})
else()
    message(FATAL_ERROR "Could not find yaml library")
endif()
