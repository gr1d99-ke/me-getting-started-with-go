package main

import "fmt"

func main()  {
	for i := 1; i <= 3; i++ {
		fmt.Printf("i = %v\n", i)

		for j := 1; j <= 3; j++ {
			fmt.Printf("j = %v\n", j)
		}
	}
}
