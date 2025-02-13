#!/usr/bin/env bash
set -e

PROJECT_ROOT="$(readlink -e $(dirname "$BASH_SOURCE[0]")/../)"
OUT_DIR=${PROJECT_ROOT}/_out

TAG="${1:-v0.0.0}"
VERSION=${TAG#v}

rm -rf ${OUT_DIR}
mkdir -p ${OUT_DIR}

cp deploy/operator.yaml ${OUT_DIR}/operator.yaml
sed -i "s/<IMAGE_VERSION>/${VERSION}/g" ${OUT_DIR}/operator.yaml

cp manifests/node-maintenance-operator/${TAG}/* ${OUT_DIR}/





