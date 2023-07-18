#!/bin/sh -xe


if [ -z "$1" ] 
then
	llvm_config=llvm-config
else
	llvm_config="$1"
fi

export CGO_CPPFLAGS="`$llvm_config --cppflags`"
export CGO_CXXFLAGS=-std=c++14
export CGO_LDFLAGS="`$llvm_config --ldflags --libs --system-libs all`"
code .