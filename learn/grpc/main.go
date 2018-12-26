package main

import (
	fmt "fmt"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("127.0.0.1")
	if err != nil {
		fmt.Println(err)
	} else {
		cd := &DeviceCountConditional{}
		c := NewWaterServiceClient(con)
		rsp, err := c.GetDeviceCount(nil, cd)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(rsp.String())
		}
	}
}
