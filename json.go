package main

import "fmt"

type Person struct {
	Name string
	Add  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v Address)", p.Name, p.Add)
}

func main() {
	a := Person{"Mathew Stevey", 42}
	fmt.Println(a)
}




