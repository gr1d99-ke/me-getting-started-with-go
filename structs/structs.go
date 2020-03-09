package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	age int
}

type Child struct {
	Person
}

type Parent struct {
	Person
}

func main()  {
	parent := Parent{
		Person{
			firstName:"aiii",
			lastName:"eeeer",
		},
	}

	fmt.Println(parent)

	var student = &Person{age:10}
	d := student
	d.age = 11
	fmt.Println(student)
	fmt.Println(d)
}
