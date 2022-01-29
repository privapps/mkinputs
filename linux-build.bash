#!/usr/bin/env bash


gobuild (){
    package=$1
    if [[ -z "$package" ]]; then
    echo "usage: $0 <package-name>"
    exit 1
    fi

    package_name=$package

    platforms=("linux/amd64" "windows/amd64")

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
        if [ $GOARCH = "arm64" ]; then # TODO not work yet
            ext_flag='CC=aarch64-linux-gnu-gcc-6'
        elif [ $GOOS = "windows" ]; then
            ext_flag='CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_LDFLAGS=-static'
            output_name+='.exe'
        else
            ext_flag='CC=gcc'
        fi
        env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=1 $ext_flag go build -o $folder/$output_name $package
        if [ $? -ne 0 ]; then
            echo 'An error has occurred! Aborting the script execution...'
            exit 1
        fi
    done
}

for i in `find cmd -name "*go"`; do
  gobuild ./$i
done