%module greet

%insert(cgo_comment) %{
#cgo CXXFLAGS: -Iinclude -Ideps/fmt/include
%}

%include <std_string.i>

%{
#include "greet.hpp"
%}

%include "greet.hpp"
