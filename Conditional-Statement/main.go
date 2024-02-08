package main

import "fmt"

// Statement = 2
const (
	first = iota
	Second
	third = 'z'
)

func main() {
	// Statement = 1
	// a := 10
	// b := 20

	// fmt.Println(test(a, b))

	fmt.Println(first, Second, third)

	// if else conditional statement

	// if rune('z') == third {
	// 	fmt.Println("Z is true")
	// } else {
	// 	fmt.Println("A is false")
	// }

	// if rune('d') == third {
	// 	fmt.Println("A is true")
	// } else {
	// 	fmt.Println("A is false")
	// }

	// if rune('c') == third {
	// 	fmt.Println("third is true")
	// } else if 0 == first {
	// 	fmt.Println("first is true")
	// } else {
	// 	fmt.Println("I'm in else case")
	// }

	var a interface{}
	a = 1.1
	switch t := a.(type) {
	case int64:
		fmt.Println("Type is an integer:", t)
	case float64:
		fmt.Println("Type is a float:", t)
	case string:
		fmt.Println("Type is a string:", t)
	case nil:
		fmt.Println("Type is nil.")
	case bool:
		fmt.Println("Type is bool:", t)
	default:
		fmt.Println("Type is unknown:")
	}

}

func test(number1 int, number2 int) (int, error) {
	return number1 + number2, nil
}
