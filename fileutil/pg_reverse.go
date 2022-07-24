package fileutil

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	// "flag"
)

func GetFullName(fileurl string) string {
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fileurl)
	return filenameWithSuffix
}

func GetExtName(fileurl string) string {
	var fileSuffix string
	fileSuffix = path.Ext(GetFullName(fileurl))
	return fileSuffix
}

func GetFileName(fileurl string) string {
	var filenameOnly string
	filenameOnly = strings.TrimSuffix(GetFullName(fileurl), GetExtName(fileurl))
	return filenameOnly
}

func ReverseFileList(fileurl string) []string {
	var dir_list []string
	fs, _ := ioutil.ReadDir(fileurl)
	for _, file := range fs {
		if file.IsDir() {
			//fmt.Println(fileurl+"/" +file.Name())
			dir_list = append(dir_list, fileurl+"/"+file.Name()+"/")
			//getFileList(fileurl+file.Name()+"/")
		} else {
			fmt.Println(fileurl + file.Name())
		}
	}
	return dir_list
}

func GetFileNoDirList(fileurl string) []string {
	var dir_list []string
	//arrlist := make([]string,0)
	err := filepath.Walk(fileurl, func(fileurl string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		} else {
			dir_list = append(dir_list, fileurl)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return dir_list
}

func GetDirList(fileurl string) []string {
	var dir_list []string
	fs, _ := ioutil.ReadDir(fileurl)
	for _, file := range fs {
		if file.IsDir() {
			dir_list = append(dir_list, fileurl+"/"+file.Name())
			// fmt.Println("==>>path:" + fileurl + file.Name())
		}
	}

	return dir_list
}

func GetDirListNoTraverse(fileurl string, maxDepth int) []string {
	var dir_list []string
	err := filepath.WalkDir(fileurl, func(fileurl string, d fs.DirEntry, err error) error {
		if d == nil {
			return err
		}
		if d.IsDir() && strings.Count(fileurl, string(os.PathSeparator)) > maxDepth {
			// dir_list = append(dir_list, fileurl)
			// fmt.Println("==>> skip", fileurl)
			return nil
		} else if d.IsDir() {
			dir_list = append(dir_list, fileurl)
		} else {
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return dir_list
}

func GetDirListTraverse(fileurl string) []string {
	var dir_list []string
	err := filepath.Walk(fileurl, func(fileurl string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			dir_list = append(dir_list, fileurl)
		} else {
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return dir_list
}

func GetFileDirList(fileurl string, sext string, bfilter bool) []string {
	var dir_list []string
	//arrlist := make([]string,0)
	err := filepath.Walk(fileurl, func(fileurl string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			if bfilter {
				if GetExtName(fileurl) == sext {
					dir_list = append(dir_list, fileurl)
				}
			} else {
				dir_list = append(dir_list, fileurl)
			}
		} else {
			return nil
			//dir_list = append(dir_list, fileurl)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return dir_list
}

// func main(){
//         flag.Parse()
//         root := flag.Arg(0)
// 		arrlist := getFilelist(root)
// 		for i := 0; i < len(arrlist); i++ {
// 			fmt.Print(arrlist[i], "\n")
// 			os.RemoveAll(arrlist[i])
// 		}
// }
