#!/usr/bin/env bash

REPOSITORY=$1
COMMIT_ID=$2

go get ${REPOSITORY}
cd $HOME/go/src/
cd ${REPOSITORY}
git checkout ${COMMIT_ID};
dep init;
dep ensure;
go test -coverprofile=cov.out -race -v $(go list ./... | grep -v /vendor/);
go tool cover -func=cov.out | grep total: | awk ' {print $3} ';
rm -rf $HOME/go/src/${REPOSITORY}
