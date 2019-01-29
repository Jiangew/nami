package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"James", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Anna", age: 25})

	s := person{name: "Sean", age: 26}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 28
	fmt.Println(sp.age)

	fmt.Println(s)
	fmt.Println(sp)
}
