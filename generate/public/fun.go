package public

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
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

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadFirstLine(filepath string) string {
	// 创建句柄
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	//func NewReader(rd io.Reader) *Reader {}，返回的是bufio.Reader结构体
	r := bufio.NewReader(fi) // 创建 Reader

	lineBytes, err := r.ReadBytes('\n')
	//去掉字符串首尾空白字符，返回字符串
	line := strings.TrimSpace(string(lineBytes))
	if err != nil && err != io.EOF {
		panic(err)
	}
	return line
}
