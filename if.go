package main

import (
	"fmt"
)

func main() {
	foo := 18

	if foo > 20 {
		fmt.Println("greater than 20")
	}

	if foo < 20 {
		fmt.Println("smaller than 20")
	}
	if foo == 20 {
		fmt.Println("20")
	}
}
