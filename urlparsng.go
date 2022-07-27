package main

import (
	"fmt"
	"net/url"
)
const myurl="https://lco.dev:3000/learn?coursename=reactjs"
func main(){
	fmt.Println("welcome tourl parsing")
	fmt.Println(myurl)
	result,_:=url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	qparams:=result.Query()
	fmt.Printf("type of query params is %T\n",qparams)
	fmt.Println(qparams["coursename"])
	for _,val:=range qparams{
		fmt.Println("value is",val)
	}
	partsofurl:=&url.URL {
		Scheme: "https",
		Host: "lco.dev", 
	}
	anotherUrl:=partsofurl.String()
	fmt.Println(anotherUrl)
}