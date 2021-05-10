#!/usr/bin/env bash

set -e

# process arguments
if [[ "$#" -ne 6 ]]; then
    echo "usage: $0 name os arch src dist lock"
    echo "given: $0 $@"
    echo "  Valid arch and os values are those supported by GOARCH and GOOS"
    exit 1
fi

name=$1
os=$2
arch=$3
src=$4
dist=$5
lock=$6

output=$dist/$name'-'$os'-'$arch
if [[ $os = "windows" ]];
then
    output+='.exe'
fi

serialized=$(sed 's/\s/\\s/g' $lock | tr -d '\n')

printf "Building %s...\n" $output
env GCO_ENABLED=0 GOOS=$os GOARCH=$arch \
go build \
    -ldflags "-X main.DependencyRequirements=$serialized" \
    -o $output $src
