#!/usr/bin/env bash

set -e

platforms=(
    "windows/amd64"
    "windows/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "linux/386"
    "linux/amd64"
    "linux/arm"
    "linux/arm64"
)

mkdir -p dist

for platform in "${platforms[@]}"; do
  export GOOS=${platform%%/*}
  export GOARCH=${platform#*/}

  exe_extension=""
  if [ "$GOOS" == "windows" ]; then
    exe_extension=".exe"
  fi

  go build -o dist/rzjd-$GOOS-$GOARCH$exe_extension cmd/*
done

tar -cf dist.tar.gz dist
