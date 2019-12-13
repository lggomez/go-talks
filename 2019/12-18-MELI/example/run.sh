#!/usr/bin/env bash
echo $BASH_VERSION

rm server_app
rm client_app
rm example_trace

echo "compiling dependencies"
./make.sh
echo

echo "starting server"
./server_app &
echo

echo "starting client"
./client_app &
echo

sleep 2
echo "collecting trace"
curl -v -XGET "localhost:5050/debug/pprof/trace?seconds=5" --output example_trace