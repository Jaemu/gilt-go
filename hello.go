package main

import (
	"fmt"
	"log"
)

var alice = Person{"Alice", Female}
var bob = Person{"Bob", Unspecified}

type SexT int

//iota!
//get incremented each time it's used
const (
	Unspecified SexT = iota
	//iota is then implied
	//we can also do math e.g. iota * 2
	Male
	Female
)

//Go declares visibility
//package-internal = lowercase
//package-external = uppercase
type Person struct {
	Name string
	Sex  SexT
}

func main() {
	//shortand for type inference
	//longhand for type declaration
	//% gives a struct,
	//v is name of fields
	fmt.Printf("%+v \n", alice)

}
