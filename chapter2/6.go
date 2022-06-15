package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	x := hello.Input("type a number")
	fmt.Print(x + "月は、")
	switch n, err := strconv.Atoi(x); n {
	case 0:
		fmt.Println("not number")
		fmt.Println(err)
	case 1, 2, 12:
		fmt.Println("winter")
	case 3, 4, 5:
		fmt.Println("spring")
	case 6, 7, 8:
		fmt.Println("summer")

	case 9, 10, 11:
		fmt.Println("atomun")
	default:
		fmt.Println("not found")

	}
}
