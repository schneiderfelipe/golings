// structs3
// Make me compile!
package main

import "fmt"
import "strings"

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FullName() string {
	b := strings.Builder{}
	if _, err := b.WriteString(p.firstName); err != nil {
		panic("unable to write first name")
	}
	if err := b.WriteByte(' '); err != nil {
		panic("unable to write space")
	}
	if _, err := b.WriteString(p.lastName); err != nil {
		panic("unable to write last name")
	}
	return b.String()
}

func main() {
	person := Person{firstName: "Maurício", lastName: "Antunes"}
	fmt.Printf("Person full name is: %s\n", person.FullName()) // here it must output Person full name is: Maurício Antunes
}
