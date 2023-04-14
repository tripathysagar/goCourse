package main

import "fmt"

var name = "Foo"
var firstName string = "hola"

const (
	Version = 1 // for exporting outside of the package
	keylen  = 1 // for using inside of the package
)

func main() {
	name := 30            // infers the int type to the variable
	var val float32 = 4.5 // explist variable type declaration.
	// if no value given default value for the data type will be assigned

	fmt.Println(name, val)
}
