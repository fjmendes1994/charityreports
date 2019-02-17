#!/usr/bin/env bash

REPOSITORY=$1
COMMIT_ID=$2

go get ${REPOSITORY} >/dev/null 2>&1
cd ${GOPATH}/src/ >/dev/null 2>&1
cd ${REPOSITORY} >/dev/null 2>&1
git checkout ${COMMIT_ID} >/dev/null 2>&1;
dep init >/dev/null 2>&1;
dep ensure >/dev/null 2>&1;
go test -coverprofile=cov.out -race -v $(go list ./... | grep -v /vendor/) >/dev/null 2>&1;
go tool cover -func=cov.out | grep total: | awk ' {print $3} ';
rm -rf ${GOPATH}/src/${REPOSITORY} >/dev/null 2>&1
