%module greet

%insert(cgo_comment) %{
#cgo CXXFLAGS: -I../include
%}

%include <std_string.i>

%{
#include "greet.hpp"
%}

%include "greet.hpp"