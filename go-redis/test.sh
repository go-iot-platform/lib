#!/usr/bin/env bash

set -e
echo "" > coverage.txt

# for d in $(go list ./internal/platform/controller/... | grep -v -e internal/mock -e cmd/api/server); do
#     go test -v -race -coverprofile=profile.out -covermode=atomic "$d"
#     if [ -f profile.out ]; then
#         cat profile.out >> coverage.txt
#         rm profile.out
#     fi
# done

for d in $(go list ./...); do
    go test -v -race -coverprofile=profile.out -covermode=atomic "$d"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done