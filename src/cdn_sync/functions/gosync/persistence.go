package gosync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"cdn_sync/drivers/aliyun"
	"cdn_sync/utils"
)

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
		fmt.Println("Read Source Dir error:", err)
	}

	for _, dirInfo := range dirInfos {
		if dirInfo.IsDir() {
			newSubDir := subDir + dirInfo.Name()
			ReadDir(filesMap, sourceDir, newSubDir)
		} else {
			filePath := currentDir + dirInfo.Name()
			fileByte, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("read file err:", err)
			}

			fileContent := string(fileByte)
			re, err := regexp.Compile("<li>Build <small>&copy; .*</small></li>")
			if err != nil {
				fmt.Println("init regexp err:", err)
			}

			fileContent = re.ReplaceAllString(fileContent, "")
			md5sum := utils.Md5sum(string(fileContent))

			filesMap[strings.Replace(filePath, sourceDir, "", 1)] = md5sum
		}
	}
}

func CacheWrite(m map[string]interface{}, cacheFile string) error {
	_, err := os.Stat(cacheFile)
	if err != nil {
		f, _ := os.Create(cacheFile)
		defer f.Close()
	}

	_j, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return err
	}

	var j bytes.Buffer
	err = json.Indent(&j, _j, "", "  ")

	err = ioutil.WriteFile(cacheFile, []byte(j.String()), 0644)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func CacheRead(filename string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(filename)
	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		fmt.Println("Unmarshal failed, ", err)
		return nil, err
	}
	return m, nil
}

func isStartWitch(str string, excludeList []string) bool {
	for _, e := range excludeList {
		if strings.HasPrefix(str, e){
			return true
		}
	}
	return false
}

func SyncLocalToOSS(aliyunOSSConfig *aliyun.AliyunOSSConfig, sourceDir, metaKey, cacheFile string, excludeList []string) error {
	if metaKey == "" {
		metaKey = "Content-Md5sum"
	}
	if cacheFile == "" {
		cacheFile = "/tmp/" + aliyunOSSConfig.BucketName + ".json"
	}
	fmt.Println("Begin to sync", sourceDir, "files, metaKey is", metaKey, ", cacheFile is", cacheFile, ", exclude file or direct is", excludeList)

	// read local files
	_filesMap := make(map[string]interface{})
	ReadDir(_filesMap, sourceDir, "")

	filesMap := make(map[string]interface{})
	for k, _ := range _filesMap {
		if isStartWitch(k, excludeList) {
			fmt.Println("Skip", k, "by exclude rule.")
			continue
		}
		filesMap[k] = _filesMap[k]
	}

	// list oss object metadata
	objectsMap := make(map[string]interface{})
	_, err := os.Stat(cacheFile)
	if err != nil {
		objectsMap, err = aliyun.ListObjects(aliyunOSSConfig, metaKey)
		if err != nil {
			aliyun.HandleError(err)
		}
	} else {
		objectsMap, err = CacheRead(cacheFile)
		if err != nil {
			aliyun.HandleError(err)
		}
	}

	// get diff map
	justM1, justM2, diffM1AndM2, err := utils.DiffMap(filesMap, objectsMap)
	if err != nil {
		aliyun.HandleError(err)
	}

	// do upload new file
	fmt.Println("Do upload new files:")
	for k, v := range justM1 {
		metasMap := make(map[string]interface{})
		metasMap[metaKey] = v
		err := aliyun.PutObjectFromFile(aliyunOSSConfig, k, sourceDir + "/" + k, metasMap)
		if err != nil {
			aliyun.HandleError(err)
			fmt.Println("Upload OSS Object", k, "Error:", err)
		}
		fmt.Println("Upload OSS Object", k, "Done.")
	}

	fmt.Println("Do update diff files:")
	for k, v := range diffM1AndM2 {
		metasMap := make(map[string]interface{})
		metasMap[metaKey] = v
		err := aliyun.PutObjectFromFile(aliyunOSSConfig, k, sourceDir + "/" + k, metasMap)
		if err != nil {
			aliyun.HandleError(err)
			fmt.Println("Update OSS Object", k, "Error:", err)
		}
		fmt.Println("Update OSS Object", k)
	}

	fmt.Println("Do delete files:")
	for k, _ := range justM2 {
		err := aliyun.DeleteObject(aliyunOSSConfig, k)
		if err != nil {
			fmt.Println("Delete OSS Object", k, "Error:", err)
		}
		fmt.Println("Delete OSS Object", k)
	}

	// cache new map to file
	_, err = os.Stat(cacheFile)
	if err == nil {
		_ = os.Truncate(cacheFile, 0)
	}

	err = CacheWrite(filesMap, cacheFile)
	if err != nil {
		fmt.Println("cache file map to file fail.")
	} else {
		fmt.Println("write cache success! path:", cacheFile)
	}

	fmt.Println("Sync done! files is", filesMap)

	return nil
}
