package main

import "fmt"

func myage[ag any](mage ag){
	fmt.Println(mage)
}

func main(){
	var age int32=24
	var age2 float32=45.7
	myage(age)
	myage(age2)

}