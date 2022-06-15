package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata{"Taro", []int{1, 2, 3}}
	hanako := Mydata{Name: "Hanako", Data: []int{9, 8, 7}}
	fmt.Println(taro)
	fmt.Println(hanako)
}
