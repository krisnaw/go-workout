package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Example usage

	fmt.Println("Hello, world!")

	p := Person{Name: "Bob", Age: 25}
	fmt.Println("inside scope", p.Name)
	p.modifyPersonName()
	fmt.Println("outside scope", p.Name)
}

func (p *Person) modifyPersonName() {
	p.Name = "John"
}
