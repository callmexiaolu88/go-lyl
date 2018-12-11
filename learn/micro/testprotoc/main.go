package main

import (
	"fmt"
	"io/ioutil"
	mathFunction "lyl/learn/micro/proto"
	"os"

	"github.com/golang/protobuf/proto"
)

const file string = "a.txt"

func main() {
	var ad = &mathFunction.AddRequest{
		P1: 10,
		P2: 12,
		P3: 11,
	}
	bytes, err := proto.Marshal(ad)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)
	ioutil.WriteFile(file, bytes, os.ModeAppend)

	//read file
	bytes, err = ioutil.ReadFile(file)

	var bb = new(mathFunction.AddResponse)
	err = proto.Unmarshal(bytes, bb)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bb)
}
