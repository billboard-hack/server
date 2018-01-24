#!/bin/bash
GOOS=windows GOARCH=amd64 go build -o wsServer_x64.exe main.go
GOOS=windows GOARCH=386 go build -o wsServer_x86.exe main.go
GOOS=darwin GOARCH=amd64 go build -o wsServer_OSX main.go
mv wsServer_x64.exe ./bin/wsServer_x64.exe
mv wsServer_x86.exe ./bin/wsServer_x86.exe
mv wsServer_OSX ./bin/wsServer_OSX