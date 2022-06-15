package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[0:2]
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 1000)
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 1000)
	fmt.Println(a)
	fmt.Println(b)
}
