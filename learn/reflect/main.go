package main

import (
	"fmt"
	"reflect"
)

type people struct {
	Name    string "this is name"
	ID      int    `address:"this is id" id:"123"`
	Age     int    `json:"age"`
	Address string
}

func (p people) Test() (string,error) {
	return p.Name,nil
}

func (p *people) Ptrname(str, str1 string) string {
	return p.Name
}

func main() {
	s := people{
		Name:    "lyl",
		ID:      101,
		Age:     10,
		Address: "shanghai"}

	tp := reflect.TypeOf(s)
	ptp := reflect.TypeOf(&s)
	tps := []reflect.Type{tp, ptp}
	for _, p := range tps {
		for index := 0; index < p.NumMethod(); index++ {
			m := p.Method(index)
			fmt.Println(m)
			fmt.Println(m.Type.NumIn())
			fmt.Println(m.Type.NumOut())
		}
		fmt.Println("====================")
	}
}
