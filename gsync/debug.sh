#!/bin/bash

set -x

xdep ensure
go mod vendor
go build -tags netgo
./gsync \
  -cdnType "aliyun" \
  -accessKeyID ${ALICLOUD_ACCESS_KEY} \
  -accessKeySecret ${ALICLOUD_SECRET_KEY} \
  -endpoint "oss-cn-hangzhou.aliyuncs.com" \
  -bucketName "dev-blog-xiexianbin-cn" \
  -sourceDir "/Users/xiexianbin/work/code/github/xiexianbin/xiexianbin.github.io/public"
