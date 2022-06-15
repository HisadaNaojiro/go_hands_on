package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fs, er := ioutil.ReadDir(".")
	if er != nil {
		panic(er)
	}

	for _, f := range fs {
		fmt.Println(f.Name(), "(", f.Size(), ")")
	}
}