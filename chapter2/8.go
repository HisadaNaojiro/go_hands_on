package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err == nil {
		println(x + "は、")
	} else {
		return
	}
	switch {
	case n%2 == 0:
		fmt.Println("偶数")
	case n%2 == 1:
		fmt.Println("奇数")
	}
}
