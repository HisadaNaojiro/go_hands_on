package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	m["total"] = m["a"] + m["b"] + m["c"]
	fmt.Println(m)
}
