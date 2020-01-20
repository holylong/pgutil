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
			lines = append(lines, fileutil.MontainFile(arrdir[i], "wget -nc -t 1 -T 3 -e \"http_proxy=http://127.0.0.1:1080\"", "")...)
			//lines = append(lines, fileutil.MontainFile(arrdir[i], "wget -nc -t 2 -T 5", "")...)
		}

		jj := 0
		for i := 0; i < len(lines); i++ {
			if i%3000 == 0{
				jj++
				//os.Mkdir(strconv.Itoa(jj), os.ModePerm)
			}
			lines[i] = lines[i] + " -O " + strconv.Itoa(jj) + "/" + strconv.Itoa(i) + ".jpg\n"
		}

		fileutil.WriteFile(lines, "http_1s", "bat", true)
	}else{
		fmt.Println("please input a root dir!")
	}
}