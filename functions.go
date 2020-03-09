package main

import "fmt"

func main()  {
	name := "gideon"
	sayHello(&name)
	fmt.Println(name)
	d, err := sayWhaat()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(d)
	anon()
	fmt.Println(*myStruct())
}

func sayHello(name *string)  {
	fmt.Printf("Hello: %v\n", *name)
	*name = "kim"
	fmt.Println(*name)
}

func sayWhaat(values ...int) (string, error) {
	fmt.Println(len(values) < 1)
	if len(values) < 1 {
		return "", fmt.Errorf("needs some values")
	}

	return "no error", nil
}

// anon
func anon()  {
	for i := 0; i < 10; i++ {
		var f func(val int) = func(val int) {
			fmt.Printf("i = %v\n", val)
		}

		f(i)
	}
}

// methods
func myStruct() *string {
	a := sampleStruct{name:"ddd"}
	return a.sample()
}

type sampleStruct struct {
	name string
}

func (s *sampleStruct) sample() *string {
	s.name = "aiiii"
	return &s.name
}