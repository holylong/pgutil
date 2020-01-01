package main

import (
	"../../dbmanager"
	"fmt"
	"encoding/json"
	"flag"
	"../../fileutil"
	"os"
)


func main()  {
	p := &dbmanager.FileInfo{}
    p.Name = "pgutil"
    p.IsDir = true
    p.FileSize = 10000
    p.LatestModify = "2499.00"
    p.FileUrl = "1"
    data, _ := json.Marshal(p)
	fmt.Println(string(data))
	if len(os.Args) > 1{
		flag.Parse()
		root := flag.Arg(0)
		arrdir := fileutil.GetFilelist(root, "zip", false)
		for i := 0; i < len(arrdir); i++ {
			fmt.Print(arrdir[i], "\t\t")
			fmt.Println(fileutil.GetFileSize(arrdir[i]), fileutil.GetFileModTime(arrdir[i]))
		}
	}else{
		fmt.Println("please input a root dir!")
	}
}