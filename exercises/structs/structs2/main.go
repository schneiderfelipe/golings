// structs2
// Make me compile!
package main

import "fmt"

type contactDetails struct {
	phone string
}

type Person struct {
	// don't just create the phone field here. embed a new struct
	name string
	age  int
	contactDetails
}

func main() {
	contactDetails := contactDetails{"+1 999 999 999"}
	person := Person{name: "John", age: 32, contactDetails: contactDetails}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
