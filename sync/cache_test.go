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
	"fmt"
	"reflect"
	"syscall"
	"testing"
)

func TestCacheWrite(t *testing.T) {
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	filePath := "./examples/dev-blog-xiexianbin-cn.js"
	m := make(map[string]interface{})
	m["a"] = "a"
	m["b"] = "b"
	m["c"] = 3

	err := CacheWrite(m, filePath)
	if err != nil {
		fmt.Println("write", m, "to filepath", filePath, "err", err)
	}
}

func TestCacheRead(t *testing.T) {
	filePath := "./examples/persistence_test.json"
	m, err := CacheRead(filePath)

	if err != nil {
		t.Skip("read from", filePath, "err", err)
		return
	}

	t.Log("m:", m)
	t.Log("m type:", reflect.TypeOf(m))
}
