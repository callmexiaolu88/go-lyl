package main

import (
	context "context"
	fmt "fmt"

	fire "honeywell.com/foxconn/fire-platform-fire-srv/proto"
)

func StartFireCalling() {
	f := fire.NewFireService("fire-side", nil)

	// fmt.Println("fire.GetDeviceCount")
	// fc1 := &fire.DeviceCountConditional{}
	// fr1, err := f.GetDeviceCount(context.TODO(), fc1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fr1.String())
	// }

	// fmt.Println("fire.GetDeviceCounts")
	// fc2 := &fire.DeviceCountsConditional{Conditionals: []*fire.DeviceCountConditional{fc1}}
	// fr2, err := f.GetDeviceCounts(context.TODO(), fc2)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fr2.String())
	// }

	fmt.Println("fire.GetDeviceValue")
	fc3 := &fire.DeviceValueConditional{}
	fr3, err := f.GetDeviceValue(context.TODO(), fc3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fr3.String())
	}

	// fmt.Println("fire.GetDeviceValues")
	// fc4 := &fire.DeviceValuesConditional{Conditionals: []*fire.DeviceValueConditional{fc3}}
	// fr4, err := f.GetDeviceValues(context.TODO(), fc4)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fr4.String())
	// }

	// fmt.Println("fire.GetDeviceDetail")
	// fc5 := &fire.DeviceDetailConditional{}
	// fr5, err := f.GetDeviceDetail(context.TODO(), fc5)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fr5.String())
	// }

	// fmt.Println("fire.GetDeviceDetails")
	// fc6 := &fire.DeviceDetailsConditional{Conditionals: []*fire.DeviceDetailConditional{fc5}}
	// fr6, err := f.GetDeviceDetails(context.TODO(), fc6)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fr6.String())
	// }
}
