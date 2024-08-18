#! /usr/bin/bash

echo "building pavus for the following platforms: windows, linux and darwin"

if [ ! -d "bin" ]; then
  mkdir bin
fi

CGO_ENABLED=0 GOOS=windows go build -ldflags='-w -s' -o bin/pavus.exe cmd/pavus/main.go
CGO_ENABLED=0 GOOS=linux go build -ldflags='-w -s' -o bin/pavus cmd/pavus/main.go
CGO_ENABLED=0 GOOS=darwin go build -ldflags='-w -s' -o bin/pavus-darwin cmd/pavus/main.go

echo "finished, look in the bin/ folder"
