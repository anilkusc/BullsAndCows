#!/bin/bash
set -e
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
GOOS=js GOARCH=wasm go build -o main.wasm main.go 
GOOS=js GOARCH=wasm go build -o game.wasm game.go
python3 -m http.server