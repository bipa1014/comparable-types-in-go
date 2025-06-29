package main

import "fmt"

func Equal[T comparable](a, b T) bool {
	return a == b
}

func main() {

	fmt.Println(Equal(5, 5))         // true
	fmt.Println(Equal("foo", "bar")) // false
	fmt.Println(Equal(3.14, 3.14))   // true
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	fmt.Println(Equal(a, b))         // crash
}
