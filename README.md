# Github Action for Sync Code to CDN

a tools sync code to cdn, like aliyun oss.

## Environment Variables

- ACCESSKEYID: CDN Access Key ID
- ACCESSKEYSECRET: CDN Access Key Secret

## How to Use

```
    - name: Sync code to CDN
      uses: x-actions/cdn-sync@release/v1
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

## Options

```
Usage of ./cdn_sync:
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
  -sourceDir string
    	the source dir public to cdn
```
