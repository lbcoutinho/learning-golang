package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	p1 := person{
		"Petter",
		"Parker",
		contactInfo{
			"email@email.com",
			12345,
		},
	}
	p1.print()

	p2 := person{firstName: "Bruce", lastName: "Banner", contactInfo: contactInfo{"email@email.com", 12345}}
	p2.print()

	var p3 person
	p3.print()
	p3.firstName = "Tony"
	p3.lastName = "Stark"
	p3.print()

	p3.updateName("Antony")
	p3.print()
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
