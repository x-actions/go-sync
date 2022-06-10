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
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/xiexianbin/golib/logger"

	"github.com/x-actions/go-sync/utils"
)

// ReadDir traverse dir
func ReadDir(filesMap map[string]interface{}, sourceDir, subDir string) {
	if strings.HasSuffix(sourceDir, "/") == false {
		sourceDir += "/"
	}

	currentDir := sourceDir
	if subDir != "" {
		if strings.HasSuffix(subDir, "/") == false {
			subDir += "/"
		}
		currentDir += subDir
	}

	dirInfos, err := ioutil.ReadDir(currentDir)
	if err != nil {
		logger.Warnf("Read Source Dir error: %s", err.Error())
	}

	for _, dirInfo := range dirInfos {
		if dirInfo.IsDir() {
			newSubDir := subDir + dirInfo.Name()
			ReadDir(filesMap, sourceDir, newSubDir)
		} else {
			filePath := currentDir + dirInfo.Name()
			fileByte, err := ioutil.ReadFile(filePath)
			if err != nil {
				logger.Warnf("read file err: %s", err.Error())
			}

			fileContent := string(fileByte)
			re, err := regexp.Compile("<li>Build <small>&copy; .*</small></li>")
			if err != nil {
				logger.Warnf("init regexp err: %s", err.Error())
			}

			fileContent = re.ReplaceAllString(fileContent, "")
			md5sum := utils.Md5sum(fileContent)

			filesMap[strings.Replace(filePath, sourceDir, "", 1)] = md5sum
		}
	}
}
