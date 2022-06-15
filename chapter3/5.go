package main

import "fmt"

func main() {
	ar := []int{1, 2, 3}
	fmt.Println(ar)
	initial(&ar)
	fmt.Println(ar)
}
func initial(ar *[]int) {
	for i := 0; i < len(*ar); i++ {
		(*ar)[i] = 0
	}
}
