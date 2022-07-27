package main

import (
	"fmt"
	"os"
	"io"
)

func main(){
	fmt.Println("hello")
	content :="hello file has been created"
	file, err :=os.Create("./mycogfile.txt")
	if err !=nil{
		panic(err)
	}
	length , err:=io.WriteString(file,content)
	if(err !=nil){
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()

}