package main

import "fmt"

func main(){
	fmt.Println("Starting to recover")
	icausepanic()
	fmt.Println("Main: post panic")
}
func icausepanic(){
	defer recovrfn()
	fmt.Println("about to panic")
	panic("func: cannot continue to execute")
}
func recovrfn(){
	if err :=recover();err !=nil {
		fmt.Println("error recovered")
	}
}