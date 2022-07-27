package main

import (
	"fmt"
	"math"
)
const c string="hello"
func main(){
	fmt.Println(c)
	const a = 32000000
	const b=3e20
	const d=b/a
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(a))
}