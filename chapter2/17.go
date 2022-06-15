package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:3]
	fmt.Println(a)
	fmt.Println(b)
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	b[1] = 200
	fmt.Println(a)
	fmt.Println(b)
}
