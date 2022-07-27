package main

import (
	"fmt"
	"time"
)

func main(){
	i:=2
	fmt.Println("write",i,"as")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	t:=time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	temp:=func (i interface{})  {
		switch t:=i.(type){
		case bool:
			fmt.Println("boolean type")
		case int:
			fmt.Println("integer type")
		default:
			fmt.Println("String type",t)
		}
	}
	temp(true)
	temp(1)
	temp("hello")
}