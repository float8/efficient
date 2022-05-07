package file

import "os"

//FileExists 文件路径
//文件/文件夹是否存在
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return  err == nil ||os.IsExist(err)
}

//DirExists 文件路径
//文件/文件夹是否存在
func  DirExists(path string) bool {
	return FileExists(path)
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