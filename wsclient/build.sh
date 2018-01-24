#!/bin/bash
GOOS=windows GOARCH=amd64 go build -o wsClient_x64.exe wsClient.go
GOOS=windows GOARCH=386 go build -o wsClient_x86.exe wsClient.go
GOOS=darwin GOARCH=amd64 go build -o wsClient_OSX wsClient.go
mv wsClient_x64.exe ./bin/wsClient_x64.exe
mv wsClient_x86.exe ./bin/wsClient_x86.exe
mv wsClient_OSX ./bin/wsClient_OSX