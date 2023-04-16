package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

// the the below func is also hard coded i.e. "_END". To parameteraize
//func Postfixer(s string) string {
//	return s + "_END"
//}

func Postfixer(end string) TransformFunc {
	return func(s string) string {
		return s + "_END"
	}
}

// Below function is dependenat of hard coded the string function
//
//	func transformString(s string) string {
//		return strings.ToUpper(s)
//	}
//
// to make the function more flexible the upgraded implementation
func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("holla", Uppercase))
	fmt.Println(transformString("holla", Postfixer("_END")))

}
