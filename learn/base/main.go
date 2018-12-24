package main

import (
	"fmt"
)

type Pe struct {
}

func (p Pe) Pint() {
	fmt.Printf("---%p\n", &p)
}

func (p *Pe) Point() {
	fmt.Printf("---%p\n", p)
}

func main() {
	var p Pe
	p.Pint()
	p.Point()

	pp := &p
	pp.Pint()
	pp.Point()
}
