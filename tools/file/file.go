package file

import "os"

// Exists 文件路径
//文件/文件夹是否存在
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return  err == nil ||os.IsExist(err)
}

//DirExists 文件路径
//文件/文件夹是否存在
func  DirExists(path string) bool {
	return Exists(path)
}

//IsDir 文件路径
//是否为文件夹
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return  fileInfo.IsDir()
}