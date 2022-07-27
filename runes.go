package main

import "fmt"

func main(){
	var s string="a string"
	fmt.Printf("value is =%v ,type is =%T\n",s,s)
	fmt.Printf("value is =%v ,type is =%T\n",s[2],s[2])
	fmt.Printf("value is =%v ,type is =%T\n",string(s[2]),string(s[2]))
	b:=[]byte(s)
	fmt.Printf("value is =%v ,type is =%T\n",b,b)
}