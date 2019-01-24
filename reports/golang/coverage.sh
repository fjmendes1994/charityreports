#!/usr/bin/env bash
dep init >/dev/null 2>/dev/null;
dep ensure >/dev/null 2>/dev/null;
go test -coverprofile=cov.out -race -v $(go list ./... | grep -v /vendor/) >/dev/null 2>/dev/null;
go tool cover -func=cov.out | grep total: | awk ' {print $3} ';
