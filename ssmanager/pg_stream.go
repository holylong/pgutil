package ssmanager

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"bufio"
	"../sysutil"
)

func ReadFile(fileurl string)string{
	b, err := ioutil.ReadFile(fileurl) // just pass the file name
    if err != nil {
        fmt.Print(err)
	}
	return string(b)
}

func ReadLines(fileurl string)[]string{
	var linelist []string
	file, err := os.Open(fileurl)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		//lineText := scanner.Text()
		linelist = append(linelist, scanner.Text())
	}
	return linelist
}

func WriteFile(fileurl,data string)error{
	f,err := os.Create(fileurl)
    defer f.Close()
    if err !=nil {
		fmt.Println(err.Error())
		return err
    } else {
        _,err=f.Write([]byte(data))
        sysutil.CheckError(err)
	}
	return nil
}