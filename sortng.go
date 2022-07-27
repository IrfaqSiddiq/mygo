package main

import (
	"fmt"
	"sort"
)

func main(){
	str:=[]string{"a","b","c"}
	sort.Strings(str)
	fmt.Println(str)

	i:=[]int{7,4,9}
	sort.Ints(i)
	fmt.Println(i)

	b:=sort.IntsAreSorted(i)
	fmt.Println(b)
}