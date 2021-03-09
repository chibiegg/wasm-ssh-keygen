#!/bin/bash

export GOOS=js
export GOARCH=wasm

mkdir -p dist/
cd golang/
go get "golang.org/x/crypto/ssh"
go build -o ../docs/ssh-keygen.wasm ssh-keygen.go
