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

// https://github.com/xiexianbin/gsync/blob/master/aliyun/oss_test.go

package object

import (
	"os"
	"testing"
)

func TestListObjects(t *testing.T) {
	client, err := NewAliyunOSSClient(
		"dev-blog-xiexianbin-cn",
		os.Getenv("ALICLOUD_ACCESS_KEY"),
		os.Getenv("ALICLOUD_SECRET_KEY"),
		"oss-cn-hangzhou.aliyuncs.com")
	if err != nil {
		t.Skip(err.Error())
		return
	}

	objects, err := client.List("Content-Md5sum")
	if err != nil {
		t.Skip(err.Error())
		return
	}

	t.Log(len(objects))
}
