package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const url="https://lco.dev"
func main()  {
	fmt.Println("web request")
	response,err:=http.Get(url)
	if(err!=nil){
		panic(err)
	}
	fmt.Printf("type: %T\n",response)
	defer response.Body.Close()
	databytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil {
		panic(err)
	}
	s:=string(databytes)
	println(s)
	
}