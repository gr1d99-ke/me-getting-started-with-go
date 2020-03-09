package main

import "fmt"

type Teacher interface {
	Teach([]byte) (int, error)
}

type SchoolTeacher struct {}

func (st SchoolTeacher) Teach(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type MyIncrementer interface {
	Increment() int
}

type JustIncrement int

func (mic *JustIncrement) Increment() int {
	*mic++
	return int(*mic)
}

func main()  {
	var t Teacher = SchoolTeacher{}
	t.Teach([]byte("Hey I am learning go!"))
	miInt := JustIncrement(0)
	var mic = &miInt
	fmt.Println(mic.Increment())
	fmt.Println(mic.Increment())
	fmt.Println(mic.Increment())
	fmt.Println(mic.Increment())
}