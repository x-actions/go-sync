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

// https://github.com/xiexianbin/gsync/blob/master/aliyun/persistence_test.go

package sync

import (
	"github.com/xiexianbin/golib/logger"
	"os"
	"testing"
)

func TestSync(t *testing.T) {
	logger.SetLogLevel(logger.DEBUG)

	sourceDir := "/Users/xiexianbin/workspace/code/github.com/xiexianbin/note/public"
	var err error

	s, err := New(
		ALIYUN,
		"oss-cn-hangzhou.aliyuncs.com",
		"dev-blog-xiexianbin-cn",
		os.Getenv("ALICLOUD_ACCESS_KEY"),
		os.Getenv("ALICLOUD_SECRET_KEY"),
		sourceDir,
		"/tmp/test.json",
		[]string{})
	if err != nil {
		t.Skipf("init sync err: %s", err)
		return
	}

	t.Log(s.Endpoint)
	//err = s.Do("Content-Md5sum")
	//if err != nil {
	//	t.Skipf("do sync occur err: %s", err)
	//}
}
