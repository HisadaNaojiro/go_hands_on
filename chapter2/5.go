package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	x := hello.Input("type a number")
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println("偶数")
		} else {
			fmt.Println("奇数")
		}
	} else {
		fmt.Print("整数ではありません")
	}
}
