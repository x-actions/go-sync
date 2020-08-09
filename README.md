# Github Action for Sync Code to CDN

a tools sync code to cdn, like aliyun oss.

## Environment Variables

- ACCESSKEYID: CDN Access Key ID
- ACCESSKEYSECRET: CDN Access Key Secret

## How to Use

```
    - name: Sync Code to CDN
      uses: x-actions/go-sync@master
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
