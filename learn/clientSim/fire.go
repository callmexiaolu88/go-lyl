package main

import (
	context "context"
	fmt "fmt"

	fire "honeywell.com/foxconn/fire-platform-fire-srv/proto"
)

func StartFireCalling() {
	f := fire.NewFireService("fire-side", nil)

	fmt.Println("fire.GetParkDeviceCount")
	fc1 := &fire.Empty{}
	fr1, err := f.GetParkDeviceCount(context.TODO(), fc1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fr1.String())
	}

	fmt.Println("fire.GetBuildDeviceCount")
	fc2 := &fire.BuildingCondition{}
	fr2, err := f.GetBuildDeviceCount(context.TODO(), fc2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fr2.String())
	}

	fmt.Println("fire.GetBuildDeviceCounts")
	fc3 := &fire.BuildingConditions{Conditionals: []*fire.BuildingCondition{fc2}}
	fr3, err := f.GetBuildsDeviceCount(context.TODO(), fc3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fr3.String())
	}

	fmt.Println("fire.GetFloorDeviceCount")
	fc4 := &fire.FloorCondition{}
	fr4, err := f.GetFloorDeviceCount(context.TODO(), fc4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fr4.String())
	}

	fmt.Println("fire.GetFloorDeviceCounts")
	fc5 := &fire.FloorsConditions{Conditionals: []*fire.FloorCondition{fc4}}
	fr5, err := f.GetFloorsDeviceCount(context.TODO(), fc5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fr5.String())
	}
}
