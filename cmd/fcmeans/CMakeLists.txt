cmake_minimum_required(VERSION 3.15)
project(fcmeans)

set(CMAKE_CXX_STANDARD 14)

include_directories(eigen)

add_executable(fcmeans
        fcm.cpp
        fcm.h
        main.cpp)
