package main

import "fmt"

func main()  {
	user := map[string]string{
		"firstName": "Gideon",
		"lastName": "Kimutai",
	}
	_, ok := user["a"]

	fmt.Println(ok)
}
