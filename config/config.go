package config

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"os"
)

const byteNumber int32 = 1024

var Config *simplejson.Json

// 读取配置文件
func LoadConfig(fileName string) {
	buf := make([]byte, byteNumber)
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	defer f.Close()
	f.Read(buf)

	Config, _ = simplejson.NewJson(buf)
}
