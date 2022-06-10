#!/bin/bash
set -e

DEBUG="${INPUT_DEBUG}"

if [[ X"$DEBUG" == X"true" ]]; then
  set -x
  DEBUG="true"
else
  DEBUG="false"
fi

CDNTYPE=${CDNTYPE:-"aliyun"}
ACCESSKEYID=${ACCESSKEYID:-""}
ACCESSKEYSECRET=${ACCESSKEYSECRET:-""}
ENDPOINT=${ENDPOINT:-""}
BUCKETNAME=${BUCKETNAME:-""}
CACHEFILE=${CACHEFILE:-""}
EXCLUDE=${EXCLUDE:-""}
SUB_DIR=${SUB_DIR:-"public"}

if test -z "${ACCESSKEYID}"; then
  echo "ACCESSKEYID is nil, skip!"
  exit -1
fi

if test -z "${ACCESSKEYSECRET}"; then
  echo "ACCESSKEYSECRET is nil, skip!"
  exit -1
fi

if test -z "${ENDPOINT}"; then
  echo "ENDPOINT is nil, skip!"
  exit -1
fi

if test -z "${BUCKETNAME}"; then
  echo "BUCKETNAME is nil, skip!"
  exit -1
fi

echo "## Check User ##################"
whoami

echo "## Check Package Version ##################"
bash --version
gsync -v

echo "## sync to cdn ##################"

gsync \
  -cdnType "${CDNTYPE}" \
  -accessKeyID "${ACCESSKEYID}" \
  -accessKeySecret "${ACCESSKEYSECRET}" \
  -endpoint "${ENDPOINT}" \
  -bucketName "${BUCKETNAME}" \
  -cacheFile "${CACHEFILE}" \
  -exclude "${EXCLUDE}" \
  -sourceDir "/github/workspace/${SUB_DIR}"

echo "## Done. ##################"
