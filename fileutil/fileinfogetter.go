package fileutil

import (
	"path/filepath"
	"os"
	"log"
	"time"
)

func GetFileSize(fileurl string)int64  {
	var result int64
    filepath.Walk(fileurl, func(tfileurl string, f os.FileInfo, err error) error {
        result = f.Size()
        return nil
    })
    return result
}

func GetFileModTime(fileurl string) int64 {
	f, err := os.Open(fileurl)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}