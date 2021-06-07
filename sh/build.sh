#!/usr/bin/env bash

set -e

# process arguments
if [[ "$#" -ne 5 ]]; then
    echo "usage: $0 name os arch src dist"
    echo "given: $0 $@"
    echo "  Valid arch and os values are those supported by GOARCH and GOOS"
    exit 1
fi

name=$1
os=$2
arch=$3
src=$4
dist=$5

output=$dist/$name'-'$os'-'$arch
if [[ $os = "windows" ]];
then
    output+='.exe'
fi

printf "Building %s...\n" $output
env GCO_ENABLED=0 GOOS=$os GOARCH=$arch \
go build -o $output $src
