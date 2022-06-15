package main

import "fmt"

func main() {
	f := func(a []string) ([]string, string) {
		return a[1:], a[0]
	}
	m := []string{
		"a",
		"b",
		"c",
	}
	s := ""
	fmt.Println(m)
	for len(m) > 0 {
		m, s = f(m)
		fmt.Println(s+" ->", m)
	}
}
