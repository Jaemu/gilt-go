package main

import (
	"fmt"
)

var alice = Person{"Alice", Female}
var bob = Person{"Bob", Unspecified}


//iota!
//get incremented each time it's used
const(
	Unspecified = iota
	//iota is then implied
	Male
	Female
)

//Go declares visibility
//package-internal = lowercase
//package-external = uppercase
type Person struct{
	Name string
	Sex int
} 

func main() {
	//shortand for type inference
	//longhand for type declaration
	//% gives a struct, 
	//v is name of fields 
	fmt.Printf("%+v \n", alice)

}