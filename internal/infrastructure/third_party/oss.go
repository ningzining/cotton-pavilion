package third_party

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type Oss interface {
	PutObject(objectName string, reader io.Reader) error
}

type OssClient struct {
	bucket *oss.Bucket
}

type OssConfig struct {
	Bucket          string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
}

func NewOssClient(cfg OssConfig) (*OssClient, error) {
	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(cfg.Bucket)
	if err != nil {
		return nil, err
	}
	return &OssClient{bucket: bucket}, nil
}

func (o OssClient) PutObject(objectName string, reader io.Reader) error {
	if err := o.bucket.PutObject(objectName, reader); err != nil {
		return err
	}
	return nil
}
