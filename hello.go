package main

import (
	"fmt"
)

var alice = &Person{"Alice", Female}
var bob = Person{"Bob", Unspecified}

type SexT int

var eve *Person

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

//Function will automatically try to find function
//even if called on pointer to Person
//Can also define function on a pointer
func (p *Person) Greet(other Person) string {
	p.Name = "Carol"
	return fmt.Sprintf("Hey, %s! it's %s! \n", other.Name, p.Name)

}

func main() {
	//shortand for type inference
	//longhand for type declaration
	//% gives a struct,
	//v is name of fields
	fmt.Printf(alice.Greet(bob))
	fmt.Printf("Name: %s\n", alice.Name)
}
