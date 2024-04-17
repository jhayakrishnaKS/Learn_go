package main

import (
	"fmt"
	"os"
)

func readFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err!=nil{
		if fileNotFound(err){
			return nil,fmt.Errorf("file %s not found",fileName)
		}
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return data,nil
}

func fileNotFound(err error) bool {
	_,ok:=err.(*os.PathError)
	return ok
}

func main(){
	fileName:="sample.txt"
	data,err:=readFile(fileName)
	if err!=nil{
		fmt.Println("Error:",err)
		return
	}

	fmt.Println("File content:", string(data))
}