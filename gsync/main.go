package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xiexianbin/gsync/aliyun"
)

var cdnType string
var endpoint string
var bucketName string
var accessKeyID string
var accessKeySecret string
var sourceDir string
var cacheFile string
var exclude string


func init() {
	flag.StringVar(&cdnType, "cdnType", "", "cdn type (aliyun)")
	flag.StringVar(&endpoint, "endpoint", "", "CDN Bucket Endpoint")
	flag.StringVar(&bucketName, "bucketName", "", "CDN Bucket Name")
	flag.StringVar(&accessKeyID, "accessKeyID", "", "CDN Access Key ID")
	flag.StringVar(&accessKeySecret, "accessKeySecret", "", "CDN Access Key Secret")
	flag.StringVar(&sourceDir, "sourceDir", "", "the source dir public to cdn")
	flag.StringVar(&cacheFile, "cacheFile", "", "the cache file path")
	flag.StringVar(&exclude, "exclude", "", "exclude file or dir, comma-separated string")

	flag.Parse()
}

func usage() {
	flag.Usage()
	os.Exit(-1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	var err error

	switch cdnType {
	case "aliyun":
		metaKey := "Content-Md5sum"
		config := &aliyun.OSSConfig{
			Endpoint:        endpoint,
			BucketName:      bucketName,
			AccessKeyID:     accessKeyID,
			AccessKeySecret: accessKeySecret,
		}
		if cacheFile == "" {
			cacheFile = "/tmp/" + config.BucketName + ".json"
		}
		excludeList := strings.Split(exclude, ",")
		err = aliyun.SyncLocalToOSS(config, sourceDir, metaKey, cacheFile, excludeList)
	}
	if err != nil {
		fmt.Println("err:", err)
	}
}
