# dingtalk client

fork from https://github.com/xiexianbin/webhooks

## Build

```
git clone https://github.com/x-actions/cdn-sync.git
cd cdn-sync
export GOPATH=`pwd`
cd src/cdn_sync
GOOS=linux GOARCH=amd64 go build -tags netgo
```

## Usage

```
./cdn_sync:
  -accessKeyID string
    	CDN Access Key ID
  -accessKeySecret string
    	CDN Access Key Secret
  -bucketName string
    	CDN Bucket Name
  -cacheFile string
    	the cache file path
  -cdntype string
    	cdn type (aliyun)
  -endpoint string
    	CDN Bucket Endpoint
  -exclude string
    	exclude file or dir, comma-separated string
  -sourceDir string
    	the source dir public to cdn
```

demo:

```
./cdn_sync \
  -cdntype "aliyun" \
  -accessKeyID "<accessKeyID>>" \
  -accessKeySecret "<accessKeySecret>>" \
  -endpoint "oss-cn-hangzhou.aliyuncs.com" \
  -bucketName "dev-blog-xiexianbin-cn" \
  -sourceDir <some-dir>
```
