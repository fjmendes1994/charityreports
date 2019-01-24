#!/usr/bin/env bash
go get $1
cd $HOME/go/src/
cd $1
git checkout $2 >/dev/null 2>/dev/null;
dep init >/dev/null 2>/dev/null;
dep ensure >/dev/null 2>/dev/null;
go test -coverprofile=cov.out -race -v $(go list ./... | grep -v /vendor/) >/dev/null 2>/dev/null;
go tool cover -func=cov.out | grep total: | awk ' {print $3} ';
rm -rf $HOME/go/src/$1
