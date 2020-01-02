package main

import (
	"fmt"
	"flag"
	"os"
	"../../fileutil"
)

func main(){
	if len(os.Args) > 1{
		flag.Parse()
		root := flag.Arg(0)
		//arrdir := fileutil.GetFilelist(root, "zip", false)
		arrdir := fileutil.ReverseFileList(root)
		for i := 0; i < len(arrdir); i++ {
			fmt.Print(arrdir[i], "\n")
			os.RemoveAll(arrdir[i])
		}
	}else{
		fmt.Println("please input a root dir!")
	}
}