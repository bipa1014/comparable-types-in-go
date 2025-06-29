package main

import (
	"fmt"
	"reflect"
)

func main() {
	m1 := map[int][]int{
		1: {1, 2},
		2: {3, 4},
	}
	m2 := map[int][]int{
		1: {1, 2},
		2: {3, 4},
	}
	// fmt.Println(m1 == m2) 		// geht nicht 👎
	fmt.Println(reflect.DeepEqual(m1, m2)) // geht 👍
}
