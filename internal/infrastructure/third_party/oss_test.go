package third_party

import (
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/config"
	"strings"
	"testing"
)

func TestUpload(t *testing.T) {
	config.LoadConfig()

	ossClient, err := NewOssClient(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if err := ossClient.PutObject("exampleDir/example.txt", strings.NewReader("大江东去浪淘尽, 千古风流人物")); err != nil {
		t.Error(err)
		return
	}
}
