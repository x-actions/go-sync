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

- Sample Use

```
    - name: Sync Code to CDN
      uses: x-actions/go-sync@main
      with:
        provider: "aliyun"
        access-key: ${{ secrets.ACCESSKEYID }}
        access-secret: ${{ secrets.ACCESSKEYSECRET }}
        endpoint: "<ENDPOINT>"
        bucket: "<BUCKETNAME>"
        source: "/github/workspace/public"
```

- Advance Use

```
    - name: Sync Code to CDN
      uses: x-actions/go-sync@main
      with:
        provider: "aliyun"
        access-key: ${{ secrets.ACCESSKEYID }}
        access-secret: ${{ secrets.ACCESSKEYSECRET }}
        endpoint: "<ENDPOINT>"
        bucket: "<BUCKETNAME>"
        cache: "<some-path/<BUCKETNAME>.json>"
        exclude: "str1,str2"  # .git,.DS_Store
        source: "/github/workspace/public"
        ignore-expr: ""  # "<li>Build <small>&copy; .*</small></li>"
        delete-objects: true
        exclude-delete-objects: "<relative-of-source-path>,<relative-of-source-file>"
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
  -provider "aliyun" \
  -access-key ${ALICLOUD_ACCESS_KEY} \
  -access-secret ${ALICLOUD_SECRET_KEY} \
  -bucket "dev-blog-xiexianbin-cn" \
  -endpoint "oss-cn-hangzhou.aliyuncs.com" \
  -source "/Users/xiexianbin/workspace/code/github.com/xiexianbin/note/public" \
  -exclude ".git,.DS_Store" \
  -ignore-expr "<li>Build <small>&copy; .*</small></li>" \
  -delete-objects=true \
  -exclude-delete-objects "images/ads/aliyun.png,images/xiexianbin.png"
```

## Others

- fork from https://github.com/xiexianbin/webhooks
- ref for https://github.com/xiexianbin/gsync
