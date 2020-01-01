package main

import (
	"fmt"
	"flag"
	"./fileutil"
)

func main(){
	flag.Parse()
    root := flag.Arg(0)
	arrdir := fileutil.GetFilelist(root, "zip", false)
	for i := 0; i < len(arrdir); i++ {
		fmt.Print(arrdir[i], "\n")
	}
}