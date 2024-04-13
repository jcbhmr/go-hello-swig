package greet

import (
	_ "github.com/jcbhmr/go-hello-swig/greet/src"
)

//go:generate swig -go -intgosize 32 -c++ -Iinclude -Ideps/fmt/include greet.i
