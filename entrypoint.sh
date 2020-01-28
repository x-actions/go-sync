#!/bin/bash
set -e

CDNTYPE=${CDNTYPE:-"aliyun"}
ACCESSKEYID=${ACCESSKEYID:-""}
ACCESSKEYSECRET=${ACCESSKEYSECRET:-""}
ENDPOINT=${ENDPOINT:-""}
BUCKETNAME=${BUCKETNAME:-""}
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

echo "## sync to cdn ##################"

cdn_sync \
  -cdntype "${CDNTYPE}" \
  -accessKeyID "${ACCESSKEYID}" \
  -accessKeySecret "${ACCESSKEYSECRET}" \
  -endpoint "${ENDPOINT}" \
  -bucketName "${BUCKETNAME}" \
  -sourceDir "/github/workspace/${SUB_DIR}"

echo "## Done. ##################"
