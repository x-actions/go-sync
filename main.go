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

var (
	provider        string
	endpoint        string
	bucketName      string
	accessKeyID     string
	accessKeySecret string
	sourceDir       string
	cacheFile       string
	exclude         string
	excludeList     []string

	debug   bool
	help    bool
	verbose bool
	version bool
)

func init() {
	flag.StringVar(&provider, "provider", "aliyun", "cdn type (only support aliyun.)")
	flag.StringVar(&endpoint, "endpoint", "oss-cn-shanghai.aliyuncs.com", "CDN Bucket Endpoint")
	flag.StringVar(&bucketName, "bucket", "", "CDN Bucket Name")
	flag.StringVar(&accessKeyID, "access-key", "", "CDN Access Key ID")
	flag.StringVar(&accessKeySecret, "access-secret", "", "CDN Access Key Secret")
	flag.StringVar(&sourceDir, "source", "", "the source dir public to cdn")
	flag.StringVar(&cacheFile, "cache", "/tmp/<bucketName>.json", "the cache file path")
	flag.StringVar(&exclude, "exclude", "", "exclude file or dir in sourceDir, comma-separated string")

	flag.BoolVar(&debug, "d", false, "Enable the debug flag to show detail log")
	flag.BoolVar(&help, "h", false, "print this help")
	flag.BoolVar(&verbose, "V", false, "be verbose, debug model")
	flag.BoolVar(&version, "v", false, "show version")

	flag.Usage = func() {
		logger.Print("Usage: gsync -h\n")
		flag.PrintDefaults()
	}

	flag.Parse()
}

func parseParams() {
	if cacheFile == "" {
		cacheFile = fmt.Sprintf("/tmp/%s.json", bucketName)
	}

	if exclude != "" {
		excludeList = strings.Split(exclude, ",")
	}

	if sourceDir == "" {
		logger.Print("source dir is empty.")
		os.Exit(1)
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
		logger.SetLogLevel(logger.DEBUG)
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	parseParams()

	var err error
	s, err := sync.New(provider, endpoint, bucketName, accessKeyID, accessKeySecret, sourceDir, cacheFile, excludeList)
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
