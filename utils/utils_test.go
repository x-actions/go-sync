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

package utils

import (
	"testing"
)

func TestMd5sum(t *testing.T) {
	if Md5sum("webhooks") == "C10F40999B74C408263F790B30E70EFE" {
		t.Log("test md5sum is ok.")
	} else {
		t.Skip("md5sum wrong")
	}
}

func TestIsStartWitch(t *testing.T) {
	excludeList := []string{".git"}
	if IsStartWitch(".git/some-file", excludeList) {
		t.Log("test exclude is ok")
	} else {
		t.Fatal("test IsStartWitch wrong.")
	}

	if IsStartWitch("path/some-file", excludeList) == false {
		t.Log("test un-exclude is ok")
	} else {
		t.Fatal("test IsStartWitch wrong.")
	}
}

func TestTrimLeftSlash(t *testing.T) {
	if TrimLeftSlash("/abc") != "abc" {
		t.Fatal("TrimLeftSlash(\"/abc\") == \"abc\" is false")
	}
}
