package main

import (
	context "context"
	fmt "fmt"
)

func main() {
	cd := &DeviceCountConditional{}
	c := NewWaterService("waterservice", nil)
	rsp, err := c.GetDeviceCount(context.TODO(), cd)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rsp.String())
	}
}
