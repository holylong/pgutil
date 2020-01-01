package dbmanager

import (
	// "fmt"
	"encoding/json"
)

type PgFileInfo struct {
    Name         string
    FileUrl      string
    FileSize     int64
    LatestModify int64
    IsDir        bool
}

func (fi *PgFileInfo) MarshalJSONInfo(name, fileurl string, latesmodify int64, dir bool, size int64) error {
	fi.Name = name
	fi.FileUrl = fileurl
	fi.LatestModify = latesmodify
	fi.IsDir = dir
	fi.FileSize = size
    return nil
}

func MarshalJSON(name, fileurl string, latesmodify int64, dir bool, size int64) []byte {
	p := &PgFileInfo{}
    p.Name = name
    p.IsDir = true
    p.FileSize = size
    p.LatestModify = latesmodify
	p.FileUrl = fileurl
	data, _ := json.Marshal(p)
    return data
}

func UnmarshalJSONFromByte(data []byte, p *PgFileInfo)error{
    return json.Unmarshal(data, p)
}

func UnmarshalJSONFromString(data string, p *PgFileInfo)error{
    return json.Unmarshal([]byte(data), p)
}