#!/bin/sh
export GOPATH=`pwd`

os=`uname`
if [ "{$os}" == "FreeBSD" ]; then
	export CC=clang
fi

go get github.com/szferi/gomdb
go test
