package main

import (
    "path/filepath"
    "os"
	"fmt"
    "path"
    "strings"
	"flag"
)

func getFullName(fileurl string)string{
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fileurl)
	return filenameWithSuffix	
}

func getExtName(fileurl string)string{
	var fileSuffix string
	fileSuffix = path.Ext(getFullName(fileurl))
	return fileSuffix
}

func getFileName(fileurl string)string{
	var filenameOnly string
	filenameOnly = strings.TrimSuffix(getFullName(fileurl), getExtName(fileurl))
	return filenameOnly	
}

func getFilelist(path string)[]string {
		var dir_list []string
		//arrlist := make([]string,0)
        err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
                if ( f == nil ) {return err}
                if f.IsDir() {
					if getExtName(path) == ".git"{
						//println(path)
						//os.RemoveAll(path)
						dir_list = append(dir_list, path)
					}
				}
				return nil
        })
        if err != nil {
                fmt.Printf("filepath.Walk() returned %v\n", err)
		}
		return dir_list
}

func main(){
        flag.Parse()
        root := flag.Arg(0)
		arrlist := getFilelist(root)
		for i := 0; i < len(arrlist); i++ {
			fmt.Print(arrlist[i], "\n")
			os.RemoveAll(arrlist[i])
		}
}