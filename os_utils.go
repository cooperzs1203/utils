/**
* @Author: Cooper
* @Date: 2019/10/15 14:20
 */

package utils

import (
	"os"
	"path/filepath"
)

// check if the dirPath directory exists
// dirPath  dir path
func CheckDirExists(dirPath string) bool {
	return checkDirOrFileExists(dirPath , true)
}

// check if the filePath file exists
// filePath file path
func CheckFileExists(filePath string) bool {
	return checkDirOrFileExists(filePath , false)
}

// check if the path directory or file exists
// path string : file or dir path
// dir bool : check dir or file
func checkDirOrFileExists(path string, dir bool) (exists bool) {
	fileInfo , _ := os.Stat(path)
	if fileInfo != nil && fileInfo.IsDir() == dir {
		exists = true
	}
	return
}

// get total size of the directory tree path , or -1  if the directory path does not exists
// dirPath string : dir path
func GetDirTreePathTotalSize(dirPath string) int64 {
	size := int64(0)
	_ = filepath.Walk(dirPath , func(subPath string, info os.FileInfo, err error) error {
		if err != nil {
			size = -1
			return nil
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	return size
}

// get the file size , if the file path file does not exists , return -1
// filePath string : file path
func GetFileSize(filePath string) int64 {
	size := int64(-1)
	fileInfo , _ := os.Stat(filePath)
	if fileInfo != nil {
		size = fileInfo.Size()
	}
	return size
}