package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type people struct {
	Name     string `json:"name,omitempty"`
	DeviceID int
	Age      int               `json:"age,omitempty"`
	Address  string            `json:"address,omitempty"`
	Hs       map[string]string `json:"hs,omitempty"`
}

type student struct {
	DeviceId int
	Age      int `json:"age,omitempty"`
}

func (p people) Test() (string, error) {
	return p.Name, nil
}

func (p *people) Ptrname(str, str1 string) string {
	return p.Name
}

func main() {
	s1 := []string{"112", "234"}
	var s []string
	s = append(s, s1...)
	fmt.Println(s1)
	s2 := []string{"456", "789"}
	combin(s1, s2)
	fmt.Println(s1)
	fmt.Println(s2)

	p := people{DeviceID: 10000111, Age: 30, Name: "lucy", Address: "shanghai pudong",
		Hs: map[string]string{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		},
	}
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes))
		s := &student{}
		err = json.Unmarshal(bytes, s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(p)
		fmt.Println(s)
	}
}

func combin(s1, s2 interface{}) {
	v1 := reflect.ValueOf(s1)
	v2 := reflect.ValueOf(s2)
	v1 = reflect.AppendSlice(v1, v2)
	fmt.Println(v1)
	p := people{}
	p1 := &p
	fmt.Println(reflect.TypeOf(p).Name(), reflect.TypeOf(p).Kind())
	fmt.Println(reflect.TypeOf(p1).Name(), reflect.TypeOf(p1).Kind())

	b := []*int{}
	fmt.Println(reflect.TypeOf(b).Elem())
	fmt.Println(reflect.TypeOf(b).Elem().Elem())
}
