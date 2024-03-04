package main

import "fmt"

// Define a type person struct
type person struct {
	name string
}

// Attach a method speak to type person using a pointer receiver (*person)
func (p *person) speak() {
	fmt.Println("Hello, my name is", p.name)
}
func (p person) blah() {
	fmt.Println("Hello, my name is", p.name)
}

// Define a type human interface
// To implicitly implement the interface, a human must have the speak method
type human interface {
	blah()
	speak()
}

// Create a function saySomething
// It takes in a human as a parameter and calls the speak method
func saySomething(h human) {
	h.speak()
	h.blah()
}

func main() {
	// Create a value of type person
	p := person{name: "John"}

	// You CAN pass a value of type *person into saySomething
	saySomething(&p) // Passing a pointer to p

	// You CANNOT pass a value of type person into saySomething
	// saySomething(p) // This will result in a compilation error
}
