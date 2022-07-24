package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/holylong/pgutil/fileutil"
	"github.com/holylong/pgutil/iocan"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("==>> tranverse files")
	var arrdsts []string
	arrdir := fileutil.GetDirList("F:/worker/stream")
	for i := 0; i < len(arrdir); i++ {
		fmt.Println(arrdir[i], i, len(arrdir))
		//进入每个文件夹，执行git remote -v
		msg, err := iocan.RunCommand(arrdir[i], "git", "remote", "-v")
		if err != nil {
			log.Println(err)
		}
		// fmt.Println(msg)
		//解析字符串
		arrs := strings.Split(msg, "\n")
		if len(arrs) > 0 {
			// fmt.Println(arrs[0])
			urls := strings.Split(arrs[0], "\t")
			if len(urls) > 1 {
				// fmt.Println(urls[1])
				dsts := strings.Split(urls[1], " ")
				if len(dsts) > 0 {
					ss := "git clone --recursive " + dsts[0]
					arrdsts = append(arrdsts, ss)
					// fmt.Println(ss)
				}
			}
		}
	}
	var (
		fileName = "F:/worker/stream/stream.bat"
		file     *os.File
		err      error
	)
	//使用追加模式打开文件
	// file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	file, err = os.OpenFile(fileName, os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Open file err =", err)
		return
	}
	defer file.Close()
	//写入文件
	for i := 0; i < len(arrdsts); i++ {
		n, err := file.Write([]byte(arrdsts[i] + "\n"))
		if err != nil {
			fmt.Println("Write file err =", err)
			return
		}
		fmt.Println("==>> n:" + string(n))
	}
}
