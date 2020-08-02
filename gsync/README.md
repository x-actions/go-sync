# gsync

- fork from https://github.com/xiexianbin/webhooks
- ref for https://github.com/xiexianbin/gsync

## Build

```
git clone https://github.com/x-actions/gsync.git
cd gsync/gsync
GOOS=linux GOARCH=amd64 go build -tags netgo
```

## Usage

```
./gsync
Usage of ./gsync:
  -accessKeyID string
    	CDN Access Key ID
  -accessKeySecret string
    	CDN Access Key Secret
  -bucketName string
    	CDN Bucket Name
  -cacheFile string
    	the cache file path
  -cdnType string
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
./gsync \
  -cdnType "aliyun" \
  -accessKeyID "<accessKeyID>>" \
  -accessKeySecret "<accessKeySecret>>" \
  -endpoint "oss-cn-hangzhou.aliyuncs.com" \
  -bucketName "dev-blog-xiexianbin-cn" \
  -sourceDir <some-dir>
```
