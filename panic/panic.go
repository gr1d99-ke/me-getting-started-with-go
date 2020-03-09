package main

import (
	"fmt"
	"log"
)

func main()  {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("Hey!!!!")
	fmt.Println("end")
}
