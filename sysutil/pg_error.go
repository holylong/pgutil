package sysutil

import (
	"log"
	"fmt"
)

func CheckError(err error)error{
	if err !=nil {
		fmt.Println(err.Error())
		log.Fatal(err)
		return err
	} 
	return nil
}