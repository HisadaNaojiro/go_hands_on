package main

import (
	"fmt"
	"reflect"
)

type General interface{}

type GData interface {
	Set(nm string, g General) GData
	Print()
}

type Ndata struct {
	Name string
	Data []int
}

func (nd *Ndata) Set(nm string, g General) GData {
	nd.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf(0)) {
		nd.Data = g.([]int)
	}
	return nd
}

func (nd *Ndata) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

type SData struct {
	Name string
	Data []string
}

func (sd *SData) Set(nm string, g General) GData {
	sd.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf("")) {
		sd.Data = g.([]string)
	}
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s [%s] *\n", sd.Name, sd.Data)
}

func main() {
	var data = []GData{}
	data = append(data, new(Ndata).Set("Taro", []int{1, 2, 3}))
	data = append(data, new(SData).Set("Jiro", []string{"hello!", "by"}))
	data = append(data, new(Ndata).Set("Hanako", 456))
	data = append(data, new(SData).Set("Sachiko", "happy?"))

	for _, ob := range data {
		ob.Print()
	}
}
