package main

import "fmt"

var mydata struct {
	Name string
	Data []int
}

func main() {
	mydata.Name = "taro"
	mydata.Data = []int{1, 2, 3}
	fmt.Println(mydata)
}
