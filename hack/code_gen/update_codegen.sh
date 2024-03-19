#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# v0.28.5 is the last version that works with go.mod
# Get last version from github releases using curl
CODEGEN_VERSION=v0.29.3
echo "Using code-generator version ${CODEGEN_VERSION}"
#brew install coreutils (mac) or sudo apt-get install realpath (linux) if you don't have it
SCRIPT_ROOT=$(realpath $(dirname ${BASH_SOURCE})/../..) 

echo "Working with directory $SCRIPT_ROOT"
# Grab code-generator version from go.sum.
CODEGEN_PKG=$(echo `go env GOPATH`"/pkg/mod/k8s.io/code-generator@${CODEGEN_VERSION}")

echo "Using ${CODEGEN_PKG}"

# code-generator does work with go.mod but makes assumptions about
# the project living in `$GOPATH/src`. To work around this and support
# any location; create a temporary directory, use this as an output
# base, and copy everything back once generated.
TEMP_DIR=$(mktemp -d)

cleanup() {
    echo "Removing ${TEMP_DIR}"
    rm -rf ${TEMP_DIR}
}
trap "cleanup" EXIT SIGINT

echo "Temporary output directory ${TEMP_DIR}"

# Ensure we can execute.
chmod +x ${CODEGEN_PKG}/generate-groups.sh ##If you get a `no file or directory` error, try `go mod download` or `go get k8s.io/code-generator@kubernetes-1.13.5`
chmod +x ${CODEGEN_PKG}/generate-internal-groups.sh

${CODEGEN_PKG}/generate-groups.sh applyconfiguration,client,deepcopy,informer,lister \
    github.com/istio-ecosystem/admiral-api/pkg/client github.com/istio-ecosystem/admiral-api/pkg/apis \
    "admiral:v1" \
    --output-base "${TEMP_DIR}" \
    --go-header-file ${SCRIPT_ROOT}/hack/code_gen/boilerplate.go.txt

# Copy everything back.
cp -r "${TEMP_DIR}/github.com/istio-ecosystem/admiral-api/." "${SCRIPT_ROOT}/"
