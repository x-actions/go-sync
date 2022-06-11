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

// https://github.com/xiexianbin/gsync/blob/master/aliyun/persistence.go

package sync

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

// CacheRead read cache file content
func CacheRead(filename string) (map[string]interface{}, error) {
	var m map[string]interface{}
	cacheBytes, err := ioutil.ReadFile(filename)
	err = json.Unmarshal(cacheBytes, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// CacheWrite write cache to file
func CacheWrite(m map[string]interface{}, cacheFile string) error {
	_, err := os.Stat(cacheFile)
	if err != nil {
		_, _ = os.Create(cacheFile)
	}

	_j, err := json.Marshal(m)
	if err != nil {
		return err
	}

	var j bytes.Buffer
	err = json.Indent(&j, _j, "", "  ")

	err = ioutil.WriteFile(cacheFile, []byte(j.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}
