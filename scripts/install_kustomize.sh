#!/usr/bin/env bash

set -euo pipefail

[ -z "${1:-}" ] && { echo "Usage: $0 <version>"; exit 1; }

version=$1

target_name=kustomize-${version}
link_path=bin/kustomize

[ -e ${link_path} ] && rm -r ${link_path}

mkdir -p bin
ln -s "${target_name}" ${link_path}

if [ ! -e bin/"${target_name}" ]; then
    os=$(go env GOOS)
    arch=$(go env GOARCH)

    url="https://github.com/kubernetes-sigs/kustomize/releases/download/v${version}/kustomize_${version}_${os}_${arch}"
    curl -L "${url}" -o bin/"${target_name}"
    chmod u+x bin/"${target_name}"
fi
