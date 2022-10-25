package number

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	<-time.After(time.Second * 10)
	fmt.Println(111)

}

func IsFileExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil { //文件不存在
		return false
	}
	if fileInfo.IsDir() { //是目录
		return false
	}
	return true
}
