package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func isEqual(A, B interface{}) bool {

	// Find out the type of A & B is Person or not
	if _, ok := A.(*Person); ok {
		if _, ok := B.(*Person); ok {
			if A.(*Person).Name == B.(*Person).Name {
				return A.(*Person).Age == B.(*Person).Age
			} else {
				return false
			}
		}
		return false
	}
	return false
}

func main() {

	Adam := NewPerson("Adam", 18)
	Adam2 := NewPerson("Adam", 18)

	if isEqual(Adam, Adam2) {
		fmt.Printf("Gleich")
	} else {
		fmt.Printf("Ungleich")
	}
}
