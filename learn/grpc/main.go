package main

import (
	context "context"
	fmt "fmt"
)

func main() {
	c1 := &DeviceCountConditional{}
	c := NewWaterService("waterservice", nil)
	r1, err := c.GetDeviceCount(context.TODO(), c1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r1.String())
	}

	c2 := &DeviceValueConditional{}
	r2, err := c.GetDeviceValue(context.TODO(), c2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r2.String())
	}

	c3 := &DeviceDetailConditional{}
	r3, err := c.GetDeviceDetail(context.TODO(), c3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r3.String())
	}
}
