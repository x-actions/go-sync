package object

// Config base Client Config
type Config struct {
	accessKeyID     string
	accessKeySecret string
	BucketName      string
	Endpoint        string
}

// IObjectClient object client interface, all object provider must implement this interface
type IObjectClient interface {
	List(metaKey string) (map[string]interface{}, error)
	PutFromFile(objectKey, filePath string, metasMap map[string]interface{}) error
	Delete(objectKey string) error
}
