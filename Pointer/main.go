package main

import "fmt"

func main() {
	//* &
	var a *int
	var b *int
	a = ptr(12)

	b = ptr(10)

	c := *a + *b

	fmt.Println("Memory address of a",&a)
	fmt.Println("Memory address of b",&b)

	fmt.Println(c)

}

func ptr(v int) *int {
	return &v
}
