package main

import (
	"fmt"
)

func main()  {
	var a int = 1
	var b *int  = &a

	fmt.Printf("%v\n", *b)
}
