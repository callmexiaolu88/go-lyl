package main

import (
	context "context"
	"encoding/json"
	fmt "fmt"

	water "honeywell.com/foxconn/fire-platform-water-srv/proto"
)

func StartWaterCalling() {
	w := water.NewWaterService("water-side", nil)

	fmt.Println("water.GetDeviceCount")
	wc1 := &water.DeviceCountConditional{}
	wr1, err := w.GetDeviceCount(context.TODO(), wc1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(wr1.String())
	}

	fmt.Println("water.GetDeviceCounts")
	wc2 := &water.DeviceCountsConditional{Conditionals: []*water.DeviceCountConditional{wc1}}
	s, _ := json.Marshal(wc2)
	fmt.Println(string(s))
	wr2, err := w.GetDeviceCounts(context.TODO(), wc2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(wr2.String())
	}

	fmt.Println("water.GetDeviceValue")
	wc3 := &water.DeviceValueConditional{}
	wr3, err := w.GetDeviceValue(context.TODO(), wc3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(wr3.String())
	}

	fmt.Println("water.GetDeviceValues")
	wc4 := &water.DeviceValuesConditional{Conditionals: []*water.DeviceValueConditional{wc3}}
	wr4, err := w.GetDeviceValues(context.TODO(), wc4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(wr4.String())
	}

	fmt.Println("water.GetDeviceDetail")
	wc5 := &water.DeviceDetailConditional{}
	wr5, err := w.GetDeviceDetail(context.TODO(), wc5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(wr5.String())
	}

	fmt.Println("water.GetDeviceDetails")
	wc6 := &water.DeviceDetailsConditional{Conditionals: []*water.DeviceDetailConditional{wc5}}
	wr6, err := w.GetDeviceDetails(context.TODO(), wc6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(wr6.String())
	}
}
