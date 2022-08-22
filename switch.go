package main

import ("fmt")

func main() {

	test := -20

	switch {
	case test > 0:
		fmt.Println("greater than zero")
		case test < 0:
		fmt.Println("smaller than zero")
	default:
		fmt.Println("zero")
	}
}
