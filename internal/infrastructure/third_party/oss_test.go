package third_party

import (
	"github.com/spf13/viper"
	"strings"
	"testing"
	"user-center/internal/infrastructure/config"
)

func TestUpload(t *testing.T) {
	config.LoadConfig()

	ossConfig := OssConfig{
		Bucket:          viper.GetString("oss.bucket"),
		Endpoint:        viper.GetString("oss.Endpoint"),
		AccessKeyID:     viper.GetString("oss.AccessKeyId"),
		AccessKeySecret: viper.GetString("oss.AccessKeySecret"),
	}
	ossClient, err := NewOssClient(ossConfig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := ossClient.PutObject("exampleDir/example.txt", strings.NewReader("大江东去浪淘尽, 千古风流人物")); err != nil {
		t.Error(err)
		return
	}
}
