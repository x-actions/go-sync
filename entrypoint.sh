#!/bin/bash
set -e

DEBUG="${INPUT_DEBUG}"

if [[ X"$DEBUG" == X"true" ]]; then
  set -x
  DEBUG="true"
else
  DEBUG="false"
fi

if test -z "${INPUT_ACCESS_KEY}"; then
  echo "ACCESS_KEY is nil, skip!"
  exit -1
fi

if test -z "${INPUT_ACCESS_SECRET}"; then
  echo "ACCESS_SECRET is nil, skip!"
  exit -1
fi

if test -z "${INPUT_ENDPOINT}"; then
  echo "ENDPOINT is nil, skip!"
  exit -1
fi

if test -z "${INPUT_BUCKET}"; then
  echo "BUCKET is nil, skip!"
  exit -1
fi

echo "## Check User ##################"
whoami

echo "## Check Package Version ##################"
bash --version
gsync -v

echo "## sync to cdn ##################"

gsync \
  -d="${DEBUG}" \
  -provider "${INPUT_PROVIDER}" \
  -access-key "${INPUT_ACCESS_KEY}" \
  -access-secret "${INPUT_ACCESS_SECRET}" \
  -endpoint "${INPUT_ENDPOINT}" \
  -bucket "${INPUT_BUCKET}" \
  -source "${INPUT_SOURCE}" \
  -cache "${INPUT_CACHE}" \
  -exclude "${INPUT_EXCLUDE}" \
  -ignore-expr "${INPUT_IGNORE_EXPR}"

echo "## Done. ##################"
