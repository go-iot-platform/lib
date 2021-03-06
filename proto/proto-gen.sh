#!/bin/bash
# Generate Micro pb files.
# alert | api | authen | email | iot | payment | policy | resource | statistics | twilio
echo "generating protoc for project $NAME"
echo $GOPATH/pkg/mod
protoc \
-I . \
-I $GOPATH/src/github.com/envoyproxy/protoc-gen-validate \
-I $GOPATH/src \
--micro_out=. \
--go_out=. \
--validate_out="lang=go:." \
$NAME/*.proto