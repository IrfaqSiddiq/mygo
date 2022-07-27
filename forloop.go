package main

import "fmt"

func main(){
	i:=1
	for i<=5{
		fmt.Println(i)
		i=i+1
	}
	for c:=0;c<=5;c++{
		if c%2==0 {
			continue
		}
		defer fmt.Println("hello")
		defer fmt.Println("world")
		fmt.Println(c)
	}
}