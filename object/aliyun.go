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

// https://github.com/xiexianbin/gsync/blob/master/aliyun/oss.go

package object

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/xiexianbin/golib/logger"
)

// AliyunOSSClient return an Aliyun oss Client
type AliyunOSSClient struct {
	BucketName string
	Client     *oss.Client
}

// NewAliyunOSSClient return Aliyun oss client
func NewAliyunOSSClient(bucketName, accessKeyID, accessKeySecret, endpoint string) (*AliyunOSSClient, error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}

	return &AliyunOSSClient{BucketName: bucketName, Client: client}, nil
}

func (c *AliyunOSSClient) List(metaKey string) (map[string]interface{}, error) {
	objectsMap := make(map[string]interface{})

	bucket, err := c.Client.Bucket(c.BucketName)
	if err != nil {
		return nil, err
	}

	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			logger.Warnf("list oss objects err: %s", err.Error())
		}
		for _, object := range lsRes.Objects {
			headers, err := bucket.GetObjectDetailedMeta(object.Key)
			if err != nil {
				logger.Warnf("get oss object %s detail metadata err: %s", object.Key, err.Error())
			}

			objectsMap[object.Key] = headers.Get("X-Oss-Meta-" + metaKey)
		}

		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}

	return objectsMap, nil
}

func (c *AliyunOSSClient) PutFromFile(objectKey, filePath string, metasMap map[string]interface{}) error {
	bucket, err := c.Client.Bucket(c.BucketName)
	if err != nil {
		return err
	}

	logger.Debugf("Begin to put objectKey: %s, filePath: %s, metasMap: %s", objectKey, filePath, metasMap)
	err = bucket.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		return err
	}

	for k, v := range metasMap {
		switch v.(type) {
		case string:
			err = bucket.SetObjectMeta(objectKey, oss.Meta(k, v.(string)))
			if err != nil {
				return err
			}
		default:
			break
		}
	}
	logger.Debugf("--> put object %s done", objectKey)

	return nil
}

func (c *AliyunOSSClient) Delete(objectKey string) error {
	var err error
	bucket, err := c.Client.Bucket(c.BucketName)
	if err != nil {
		return err
	}

	err = bucket.DeleteObject(objectKey)
	if err != nil {
		return err
	}

	return nil
}
