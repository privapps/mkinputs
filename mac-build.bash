#!/usr/bin/env bash


gobuild (){
    package=$1
    if [[ -z "$package" ]]; then
    echo "usage: $0 <package-name>"
    exit 1
    fi

    package_name=$package

    platforms=("darwin/arm64" "darwin/amd64" "linux/arm64")

    for platform in "${platforms[@]}"
    do
        CGO_ENABLED=1
        CC=gcc
        platform_split=(${platform//\// })
        GOOS=${platform_split[0]}
        GOARCH=${platform_split[1]}
        folder=bin/$GOOS'-'$GOARCH
        mkdir -p $folder
        output_name=$(echo ${package_name} | awk -F'/' '{print $3}')

        env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=1 go build -o $folder/$output_name $package
        if [ $? -ne 0 ]; then
            echo 'An error has occurred! Aborting the script execution...'
            exit 1
        fi
    done
}

for i in `find cmd -name "*go"`; do
  gobuild ./$i
done