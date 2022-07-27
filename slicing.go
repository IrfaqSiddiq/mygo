package main

import "fmt"

func main(){
	fruit :=make([]string,3)
	fmt.Println("abc: ",fruit)
	fruit[0]="Apple"
	fmt.Println(fruit)
	fruit[1]="mango"
	fruit[2]="banana"
	fmt.Println(fruit)
	fmt.Println(len(fruit))
	a:=make([]string,len(fruit))
	copy(a,fruit)
	fmt.Println(a)
	a=append(a, "club")
	fmt.Println(a)
	l:=a[2:3]
	fmt.Println(l)
}