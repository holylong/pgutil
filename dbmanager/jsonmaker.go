package dbmanager

import (
	// "fmt"
	"encoding/json"
)

type FileInfo struct {
    Name         string
    FileUrl      string
    FileSize     int
    LatestModify string
    IsDir        bool
}

func (fi *FileInfo) UnmarshalJSON(name, fileurl, latesmodify string, dir bool, size int) error {
	fi.Name = name
	fi.FileUrl = fileurl
	fi.LatestModify = latesmodify
	fi.IsDir = dir
	fi.FileSize = size
    return nil
}

func MarshalJSON(name, fileurl, latesmodify string, dir bool, size int) []byte {
	p := &FileInfo{}
    p.Name = "pgutil"
    p.IsDir = true
    p.FileSize = 10000
    p.LatestModify = "2499.00"
	p.FileUrl = fileurl
	data, _ := json.Marshal(p)
    return data
}