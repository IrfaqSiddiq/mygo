package main

import (
	"fmt"
	"time"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	if x := 9; x < 0 {
		fmt.Println(x, "negative number")
	} else if x < 10 {
		fmt.Println(x, "is one digit number")
	}

	switch time.Now().Weekday() {

	}
}
