package main

import "fmt"

type Person struct {
	Name string
	Address int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v Address)", p.Name, p.Address)
}

func main() {
	a := Person{"Mathew Stevey", 42}
	fmt.Println(a)
}




