package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
	// "strings"
	"../../fileutil"
)

func main(){
	if len(os.Args) > 1{
		flag.Parse()
		root := flag.Arg(0)
		// arrdir := fileutil.GetFileDirList(root, ".git", false)
		// //arrdir := fileutil.ReverseFileList(root)
		// for i := 0; i < len(arrdir); i++ {
		// 	fmt.Print(arrdir[i], "\n")
		// 	//os.RemoveAll(arrdir[i])
		// }

		/*获取所有文本文件*/
		arrdir := fileutil.GetFileNoDirList(root)
		var lines []string
		for i := 0; i < len(arrdir); i++ {
			//fmt.Print(arrdir[i], "\n")
			lines = append(lines, fileutil.MontainFile(arrdir[i], "wget -e 'https_proxy=http://127.0.0.1:1080'", "")...)
		}

		for i := 0; i < len(lines); i++ {
			lines[i] = lines[i] + " -O " + strconv.Itoa(i) + ".jpg\n"
		}

		fileutil.WriteFile(lines, "abs", "bat", true)
	}else{
		fmt.Println("please input a root dir!")
	}
}