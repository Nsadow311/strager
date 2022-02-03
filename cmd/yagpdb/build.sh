#!/bin/bash
VERSION=$(git describe --tags)
echo Building version $VERSION
go build -ldflags "-X github.com//Nsadow311/stranger/common.VERSION=${VERSION}"