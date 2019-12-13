#!/usr/bin/env bash

go_version=`go version`
echo "$go_version"
echo

(cd ./client/main && go build)
mv ./client/main/client client_app

(cd ./server/main && go build)
mv ./server/main/server server_app