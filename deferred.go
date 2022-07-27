package main

import "fmt"

func main(){
	fmt.Println("hello")
	defer fmt.Println("world")
	myDefer()
}
func myDefer(){
	for i := 0; i < 5; i++{
		defer fmt.Println(i)
	}
}