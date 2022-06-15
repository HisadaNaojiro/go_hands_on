package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata{"Taro", []int{1, 2, 3}}
	fmt.Println(taro)
	rev(&taro)
	fmt.Println(taro)
}

func rev(md *Mydata) {
	od := (*md).Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}
