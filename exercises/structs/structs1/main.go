// structs1
// Make me compile!
package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func main() {
	person := Person{"Felipe", 33}
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
