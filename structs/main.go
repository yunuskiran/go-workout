package main

import "fmt"

type contactInfo struct {
	email   string
	address string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	me := person{
		firstName: "Yunus",
		lastName:  "KIRAN",
		contactInfo: contactInfo{
			email:   "yunuskiran54@gmail.com",
			address: "5 Nova Ct",
		},
	}

	me.updateName("Esra")
	me.print()
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
