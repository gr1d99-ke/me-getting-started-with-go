package main

import "fmt"

func main()  {
	var w writer = ConsoleWriter{}
	w.Write([]byte("hello there"))
	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	inc.Increment()
	fmt.Println(inc.Increment())
}

type writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println("Hey")

	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
