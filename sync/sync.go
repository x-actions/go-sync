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
	"fmt"
	"github.com/x-actions/go-sync/object"
	"github.com/x-actions/go-sync/utils"
	"os"
	"path"
	"sync"

	"github.com/xiexianbin/golib/concurrentmap"
	"github.com/xiexianbin/golib/logger"
	xutils "github.com/xiexianbin/golib/utils"
)

type Sync struct {
	Provider                 string
	Endpoint                 string
	BucketName               string
	accessKeyID              string
	accessKeySecret          string
	SourceDir                string
	CacheFile                string
	ExcludeList              []string
	IgnoreExprList           []string
	DeleteObjects            bool
	ExcludeDeleteObjectsList []string

	// api which implement object.IObjectClient
	ObjectAPI object.IObjectClient
}

// uploadFiles upload new/exist files to object provider
func (s *Sync) uploadFiles(m map[string]interface{}, metaKey, sourceDir, action string, ch chan bool) {
	if action == "" {
		action = "new"
	}

	// new ConcurrentMap
	cMap := concurrentmap.New()

	// set key and value to cMap
	for k, v := range m {
		cMap.Set(k, v)
	}

	// new goroutine
	wg := sync.WaitGroup{}
	wg.Add(concurrentmap.ShareCount)

	logger.Infof("Do upload %s files:", action)
	for i := 0; i < concurrentmap.ShareCount; i++ {
		// pre shared map, new goroutine to statics
		go func(ms *concurrentmap.Shared, index int) {
			count := 1
			sum := len(ms.Items)
			ms.Mu.RLock() // read locak
			for k := range ms.Items {
				process := fmt.Sprintf("(map shard id %d, [%d/%d])", index, count, sum)
				metasMap := make(map[string]interface{})
				v, _ := cMap.Get(k)
				metasMap[metaKey] = v
				err := s.ObjectAPI.PutFromFile(k, path.Join(sourceDir, k), metasMap)
				if err != nil {
					logger.Warnf("Upload %s OSS Object %s %s err: %s", action, process, k, err.Error())
				} else {
					logger.Debugf("Upload %s OSS Object %s %s Done", action, process, k)
				}
				count++
			}
			ms.Mu.RUnlock() // unlock
			wg.Done()
		}((*cMap)[i], i)
	}

	// wait all goroutine stop
	wg.Wait()
	ch <- true
	logger.Infof("upload %s files Done", action)
}

// syncDelFiles Do delete files from aliyun oss
func (s *Sync) deleteFiles(m map[string]interface{}, ch chan bool) {
	if s.DeleteObjects == false {
		ch <- true
		logger.Infof("skip delete %d objects by delete-objects==false", len(m))
		return
	}

	// new ConcurrentMap
	cMap := concurrentmap.New()

	skipCount := 0
	// set key and value to cMap
	for k, v := range m {
		// skip delete special objects
		if utils.IsStartWitch(k, s.ExcludeDeleteObjectsList) {
			logger.Debugf("skip delete object %s by exclude-delete-objects rule", k)
			skipCount++
			continue
		}

		cMap.Set(k, v)
	}

	// new goroutine
	wg := sync.WaitGroup{}
	wg.Add(concurrentmap.ShareCount)

	logger.Info("Do delete files:")
	for i := 0; i < concurrentmap.ShareCount; i++ {
		// pre shared map, new goroutine to statics
		go func(ms *concurrentmap.Shared, index int) {
			count := 1
			sum := len(ms.Items)
			ms.Mu.RLock() // read locak
			for k := range ms.Items {
				process := fmt.Sprintf("(map shard id %d, [%d/%d])", index, count, sum)
				err := s.ObjectAPI.Delete(k)
				if err != nil {
					logger.Warnf("Delete OSS Object %s %s err: %s", process, k, err)
				} else {
					logger.Debugf("Delete OSS Object %s %s Done", process, k)
				}
				count++
			}
			ms.Mu.RUnlock() // unlock
			wg.Done()
		}((*cMap)[i], i)
	}

	// wait goroutine done
	wg.Wait()
	ch <- true

	logger.Infof("delete %d files Done", len(m)-skipCount)
}

// Do do sync logic
func (s *Sync) Do(metaKey string) error {
	logger.Infof("Begin to sync source files from: %s to %s:%s, metaKey is %s, cacheFile is %s, "+
		"exclude file or direct is %s, delete-object is %v, exclude-delete-objects is %s",
		s.SourceDir, s.Provider, s.BucketName, metaKey, s.CacheFile, s.ExcludeList, s.DeleteObjects, s.ExcludeDeleteObjectsList)

	// read local files
	_filesMap := make(map[string]interface{})
	ReadDir(_filesMap, s.SourceDir, "", s.IgnoreExprList)

	filesMap := make(map[string]interface{})
	for k := range _filesMap {
		if utils.IsStartWitch(k, s.ExcludeList) {
			logger.Debugf("skip sync %s by exclude rule", k)
			continue
		}
		filesMap[k] = _filesMap[k]
	}

	// list oss object metadata
	objectsMap := make(map[string]interface{})
	_, err := os.Stat(s.CacheFile)
	if err != nil {
		objectsMap, err = s.ObjectAPI.List(metaKey)
		if err != nil {
			logger.Warnf(err.Error())
		}
	} else {
		objectsMap, err = CacheRead(s.CacheFile)
		if err != nil {
			logger.Warnf("read cache file err: %s", err.Error())
		}
	}

	// get diff map
	justM1, justM2, diffM1AndM2 := xutils.DiffMap(filesMap, objectsMap)
	if err != nil {
		logger.Warnf("read cache file err: %s", err.Error())
	}

	// signal channel
	newFileChan, diffFileChan, deleteFileChan := make(chan bool), make(chan bool), make(chan bool)

	// do upload new file
	go s.uploadFiles(justM1, metaKey, s.SourceDir, "new", newFileChan)

	// do update diff files
	go s.uploadFiles(diffM1AndM2, metaKey, s.SourceDir, "update", diffFileChan)

	// do delete files
	go s.deleteFiles(justM2, deleteFileChan)

	<-newFileChan
	<-diffFileChan
	<-deleteFileChan

	// cache new map to file
	_, err = os.Stat(s.CacheFile)
	if err == nil {
		_ = os.Truncate(s.CacheFile, 0)
	}

	err = CacheWrite(filesMap, s.CacheFile)
	if err != nil {
		logger.Errorf("cache file map to file fail, err: %s", err.Error())
	} else {
		logger.Debugf("write cache success! path: %s", s.CacheFile)
	}

	logger.Infof("Sync done!")
	return nil
}

// New return new sync client
func New(provider, endpoint, bucketName, accessKeyID, accessKeySecret, sourceDir, cacheFile string,
	excludeList, ignoreExprList []string, deleteObjects bool, excludeDeleteObjectsList []string) (*Sync, error) {
	var client object.IObjectClient
	var err error

	switch provider {
	case ALIYUN:
		client, err = object.NewAliyunOSSClient(bucketName, accessKeyID, accessKeySecret, endpoint)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("un-support cdn type: %s", provider)
	}

	return &Sync{
		Provider:                 provider,
		Endpoint:                 endpoint,
		BucketName:               bucketName,
		accessKeyID:              accessKeyID,
		accessKeySecret:          accessKeySecret,
		SourceDir:                sourceDir,
		CacheFile:                cacheFile,
		ExcludeList:              excludeList,
		IgnoreExprList:           ignoreExprList,
		DeleteObjects:            deleteObjects,
		ExcludeDeleteObjectsList: excludeDeleteObjectsList,

		ObjectAPI: client,
	}, nil
}
