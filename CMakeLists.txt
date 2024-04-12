cmake_minimum_required(VERSION 3.28)
project(greet)

# add_library(greet STATIC src/greet.cpp)

include(FindSWIG)
find_package(SWIG 4.0 COMPONENTS go)
if(NOT SWIG_FOUND)
    message(FATAL_ERROR "swig 4.0 go not found")
endif()
include(UseSWIG)
swig_add_library(greet
    LANGUAGE go
    SOURCES greet/greet.i)
set_property(SOURCE greet/greet.i PROPERTY CPLUSPLUS ON)
set_property(SOURCE greet/greet.i PROPERTY COMPILE_OPTIONS "-intgosize 32")
target_include_directories(greet PRIVATE include)
target_compile_features(greet PRIVATE cxx_std_20)
