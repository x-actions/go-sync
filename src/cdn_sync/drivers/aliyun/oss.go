package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunOSSConfig struct {
	Endpoint        string
	BucketName      string
	AccessKeyID     string
	AccessKeySecret string
}

func HandleError(err error) {
	fmt.Println("Error:", err)
}

func ListObjects(aliyunOSSConfig *AliyunOSSConfig, metaKey string) (map[string]interface{}, error) {
	objectsMap := make(map[string]interface{})
	client, err := oss.New(aliyunOSSConfig.Endpoint, aliyunOSSConfig.AccessKeyID, aliyunOSSConfig.AccessKeySecret)
	if err != nil {
		HandleError(err)
		return nil, err
	}

	bucket, err := client.Bucket(aliyunOSSConfig.BucketName)
	if err != nil {
		HandleError(err)
		return nil, err
	}

	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			HandleError(err)
		}
		for _, object := range lsRes.Objects {
			headers, err := bucket.GetObjectDetailedMeta(object.Key)
			if err != nil {
				HandleError(err)
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

func PutObjectFromFile(aliyunOSSConfig *AliyunOSSConfig, objectKey, filePath string, metasMap map[string]interface{}) error {
	client, err := oss.New(aliyunOSSConfig.Endpoint, aliyunOSSConfig.AccessKeyID, aliyunOSSConfig.AccessKeySecret)
	if err != nil {
		HandleError(err)
		return err
	}

	bucket, err := client.Bucket(aliyunOSSConfig.BucketName)
	if err != nil {
		HandleError(err)
		return err
	}

	fmt.Println("Begin to put objectKey:", objectKey, "filePath:", filePath, "metasMap:", metasMap)
	err = bucket.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		HandleError(err)
		return err
	}

	for k, v := range metasMap {
		fmt.Println("k:", k, "v:", v)
		switch v.(type) {
		case string:
			err = bucket.SetObjectMeta(objectKey, oss.Meta(k, v.(string)))
			if err != nil {
				HandleError(err)
				return err
			}
		default:
			break
		}
	}
	fmt.Println("--> put object", objectKey, "done.")

	return nil
}

func DeleteObject(aliyunOSSConfig *AliyunOSSConfig, objectKey string) error {
	client, err := oss.New(aliyunOSSConfig.Endpoint, aliyunOSSConfig.AccessKeyID, aliyunOSSConfig.AccessKeySecret)
	if err != nil {
		HandleError(err)
		return err
	}

	bucket, err := client.Bucket(aliyunOSSConfig.BucketName)
	if err != nil {
		HandleError(err)
		return err
	}

	err = bucket.DeleteObject(objectKey)
	if err != nil {
		HandleError(err)
		return err
	}

	return nil
}
