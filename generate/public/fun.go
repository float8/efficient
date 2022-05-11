package public

import (
	"log"
	"os"
)

func WriteFile(path string, s string) {
	_, err := os.Stat(path)
	if err == nil {
		return
	}
	f, err := os.Create(path)
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	// 关闭文件
	defer f.Close()
	_, err = f.WriteString(s)
	if err != nil {
		log.Println(err)
		return
	}
}