// Copyright 2020 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"github.com/x-actions/go-sync/sync"
	"github.com/xiexianbin/golib/logger"
	"os"
	"strings"
)

const (
	defaultCacheFile = "/tmp/<bucketName>.json"
)

var (
	provider       string
	endpoint       string
	bucket         string
	accessKey      string
	accessSecret   string
	source         string
	cacheFile      string
	exclude        string
	excludeList    []string
	ignoreExpr     string
	ignoreExprList []string

	debug   bool
	help    bool
	verbose bool
	version bool
)

func init() {
	flag.StringVar(&provider, "provider", "aliyun", "cdn type (only support aliyun.)")
	flag.StringVar(&endpoint, "endpoint", "oss-cn-shanghai.aliyuncs.com", "CDN Bucket Endpoint")
	flag.StringVar(&bucket, "bucket", "", "CDN Bucket Name")
	flag.StringVar(&accessKey, "access-key", "", "CDN Access Key ID")
	flag.StringVar(&accessSecret, "access-secret", "", "CDN Access Key Secret")
	flag.StringVar(&source, "source", "", "the source dir public to cdn")
	flag.StringVar(&cacheFile, "cache", defaultCacheFile, "the cache file path")
	flag.StringVar(&exclude, "exclude", "", "exclude file or dir in sourceDir, comma-separated string")
	flag.StringVar(&ignoreExpr, "ignore-expr", "",
		"ignore expression string, comma-separated string, replace by empty string when calculate md5 summary")

	flag.BoolVar(&debug, "d", false, "Enable the debug flag to show detail log")
	flag.BoolVar(&help, "h", false, "print this help")
	flag.BoolVar(&verbose, "V", false, "be verbose, debug model")
	flag.BoolVar(&version, "v", false, "show version")

	flag.Usage = func() {
		logger.Print("Usage: gsync -d=true\n")
		flag.PrintDefaults()
	}

	flag.Parse()
}

func parseParams() {
	if accessKey == "" || accessSecret == "" {
		logger.Print("access-key or access-secret is empty.")
		os.Exit(1)
	}

	if bucket == "" {
		logger.Print("bucket mame is empty.")
		os.Exit(1)
	}

	if source == "" {
		logger.Print("source dir is empty.")
		os.Exit(1)
	}

	if cacheFile == defaultCacheFile {
		cacheFile = fmt.Sprintf("/tmp/%s.json", bucket)
	}

	if exclude != "" {
		excludeList = strings.Split(exclude, ",")
	} else {
		excludeList = []string{}
	}

	if ignoreExpr != "" {
		ignoreExprList = strings.Split(ignoreExpr, ",")
	} else {
		ignoreExprList = []string{}
	}

}

func main() {
	if help == true {
		flag.Usage()
		return
	}

	if version == true {
		showVersion()
		return
	}

	if verbose == true || debug == true {
		logger.Print("run with Debug model...")
		logger.SetLogLevel(logger.DEBUG)
	}

	parseParams()

	var err error
	s, err := sync.New(provider, endpoint, bucket, accessKey, accessSecret, source, cacheFile, excludeList, ignoreExprList)
	if err != nil {
		logger.Errorf("init sync err: %s", err)
		os.Exit(1)
	}
	err = s.Do("Content-Md5sum")
	if err != nil {
		logger.Errorf("do sync occur err: %s", err)
		os.Exit(1)
	}
}
