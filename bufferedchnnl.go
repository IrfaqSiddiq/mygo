package main

import "fmt"

func main(){
	message :=make(chan string ,2)

	func ()  {
		message <-"Buffer"
		message <-"channel"
	}()
	fmt.Println(<-message)
	fmt.Println(<-message)
}