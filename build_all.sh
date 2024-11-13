#!/bin/bash

traverse_and_build() {
    local dir="$1"
    cd "$dir" || return
    if ls *.go &> /dev/null; then
        echo "found go files in $dir, calling go build..."
        go build
    fi
    for subdir in */; do
        if [ -d "$subdir" ]; then
            traverse_and_build "$dir/$subdir"
        fi
    done
    cd ..
}

traverse_and_build "$(pwd)"

