package main

import (
	"fmt"
	s "strings"
)
func main(){
	fmt.Println(s.Contains("test","es"))
	fmt.Println(s.Count("test","e"))
	fmt.Println(s.HasPrefix("test","t"))
	fmt.Println(s.Index("test","e"))
	fmt.Println(s.Replace("food","o","0",-1))
	fmt.Println(s.Split("a-b-c-d-e","-"))
}