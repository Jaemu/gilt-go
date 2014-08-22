package main

import (
	"fmt"
)

//Go declares visibility
//package-internal = lowercase
//package-external = uppercase
type Person struct{
	Name string
} 

func main() {

	a:= Person{"Alice"}
	//% gives a struct, 
	//v is name of fields 
	fmt.Printf("%+v \n", a)

}