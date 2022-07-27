package main

import "os"

func main()  {
	panic("problem is there")

	_,err := os.Create("/temp/file")
	if err !=nil{
		panic(err)
	}

}