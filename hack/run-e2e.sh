#!/bin/bash

# Copyright 2022 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e
set -x

FOCUS=$1

# start at the root of the repo
REPO_ROOT="$(git rev-parse --show-toplevel)"

BASEDIR="${REPO_ROOT}/tests/e2e"

BINDIR="${BASEDIR}/.build/"
mkdir -p "${BINDIR}/"

pushd "${BASEDIR}"
    go build -o "${BINDIR}/kubetest2-kops" ./kubetest2-kops
    go build -o "${BINDIR}/kubetest2-tester-kops" ./kubetest2-tester-kops
popd

export PATH="${BINDIR}/:$PATH"

kubetest2-kops --test=kops -- --test-args="-test.timeout=60m -num-nodes=0 --ginkgo.focus=${FOCUS}"
