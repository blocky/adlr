#!/usr/bin/env bash

set -e

output_name(){
    name=$1
    dist=$2
    output=$dist/$name'-'$os'-'$arch

    if [[ $os = "windows" ]];
    then
        output+='.exe'
    fi
    echo $output
}

serialize_licenselock(){
    # ldflags cannot have newlines or spaces
    licenselock=$1
    sed 's/\s/\\s/g' $licenselock | tr -d '\n'
}

# process arguments
if [[ "$#" -ne 6 ]]; then
    echo "usage: $0 name os arch src dist depreq"
    echo "given: $0 $@"
    echo "  Valid arch and os values are those supported by GOARCH and GOOS"
    exit 1
fi

name=$1
os=$2
arch=$3
src=$4
dist=$5
depreq=$6

output=$(output_name $name $dist)

serialized=$(serialize_licenselock $6)
depreq_flag="-X main.DependencyRequirements=$serialized"

printf "Building %s...\n" $output
env GCO_ENABLED=0 GOOS=$os GOARCH=$arch \
go build \
    -ldflags "$depreq_flag" \
    -o $output $src
