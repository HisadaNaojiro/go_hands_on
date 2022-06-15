package main

import "fmt"

func main() {
	data := "new"
	m1 := modify(data)
	data = "add"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())
}

func modify(d string) func() []string {
	m := []string{
		"a",
		"b",
	}
	return func() []string {
		return append(m, d)
	}
}
