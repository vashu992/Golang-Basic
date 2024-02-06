package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	var b = false
	c := 20 //:= shorthand method

	fmt.Printf("a is type of %v and b is type of %v, and c is type of %v",
		reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))

	test()
}

func test() {
fmt.Println("hello team")
}
