package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func (m Mydata) PrintData() {
	fmt.Println("myData")
	fmt.Println(m.Name)
	fmt.Println(m.Data)
	fmt.Println("end")
}

func main() {
	taro := Mydata{
		"Hanako", []int{1, 2, 3, 4, 5},
	}
	taro.PrintData()
}
