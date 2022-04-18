#!/bin/sh -xe

export CGO_CPPFLAGS="`llvm-config --cppflags`"
export CGO_CXXFLAGS=-std=c++14
export CGO_LDFLAGS="`llvm-config --ldflags --libs --system-libs all`"
go build -tags byollvm