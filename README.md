# Github Action for Sync Code to CDN

[![build-test](https://github.com/xiexianbin/go-actions-demo/actions/workflows/workflow.yaml/badge.svg)](https://github.com/xiexianbin/go-actions-demo/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/xiexianbin/go-actions-demo?status.svg)](https://pkg.go.dev/github.com/xiexianbin/go-actions-demo)
[![Go Report Card](https://goreportcard.com/badge/github.com/xiexianbin/go-actions-demo)](https://goreportcard.com/report/github.com/xiexianbin/go-actions-demo)

a tools sync code to cdn, like aliyun oss.

## Environment Variables

- ACCESSKEYID: CDN Access Key ID
- ACCESSKEYSECRET: CDN Access Key Secret

## Usage

### Use as Github Action

```
    - name: Sync Code to CDN
      uses: x-actions/go-sync@main
      env:
        CDNTYPE: "aliyun"
        ACCESSKEYID: ${{ secrets.ACCESSKEYID }}
        ACCESSKEYSECRET: ${{ secrets.ACCESSKEYSECRET }}
        ENDPOINT: "<ENDPOINT>"
        BUCKETNAME: "<BUCKETNAME>"
        CACHEFILE: "<some-path/<BUCKETNAME>.json>"
        EXCLUDE: "str1,str2"
        SUB_DIR: "public"
```

### Usage as command line

- download

```
curl -Lfs -o main https://github.com/xiexianbin/go-actions-demo/releases/latest/download/gsync-{linux|darwin|windows}
chmod +x gsync
./gsync -h
```

- or build

```
git clone https://github.com/x-actions/go-sync.git
make all
```

- usage

```
./gsync \
  -cdnType "aliyun" \
  -accessKeyID ${ALICLOUD_ACCESS_KEY} \
  -accessKeySecret ${ALICLOUD_SECRET_KEY} \
  -endpoint "oss-cn-hangzhou.aliyuncs.com" \
  -bucketName "dev-blog-xiexianbin-cn" \
  -sourceDir "/Users/xiexianbin/workspace/code/github.com/xiexianbin/note/public"
```

## Others

- fork from https://github.com/xiexianbin/webhooks
- ref for https://github.com/xiexianbin/gsync
