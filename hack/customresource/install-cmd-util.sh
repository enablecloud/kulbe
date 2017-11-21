#!/bin/bash
API_PORT=${API_PORT:-8080}
SECURE_API_PORT=${SECURE_API_PORT:-6443}
run_crd() {
  set -o nounset
  set -o errexit

  kube::log::status "Testing kubectl crd"
  kubectl "${kube_flags_with_token[@]}" create -f - << __EOF__
{
  "kind": "CustomResourceDefinition",
  "apiVersion": "apiextensions.k8s.io/v1beta1",
  "metadata": {
    "name": "application.kulbe.enablecloud.github.com"
  },
  "spec": {
    "group": "kulbe.enablecloud.github.com",
    "version": "v1",
    "names": {
      "plural": "kapps",
      "kind": "kapp"
    }
  }
}
__EOF__

kube_flags=(
    -s "http://127.0.0.1:${API_PORT}"
  )

 # token defined in hack/testdata/auth-tokens.csv
  kube_flags_with_token=(
    -s "https://127.0.0.1:${SECURE_API_PORT}" --token=admin-token --insecure-skip-tls-verify=true
  )

  if [[ -z "${ALLOW_SKEW:-}" ]]; then
    kube_flags+=("--match-server-version")
    kube_flags_with_token+=("--match-server-version")
  fi
