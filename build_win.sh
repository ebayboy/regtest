#!/bin/bash -x

FILE=$1

# ./build_win.sh service

if [ "$FILE" == "" ];then
    echo "file nil"
    exit 1
fi

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${FILE}.exe ${FILE}.go

