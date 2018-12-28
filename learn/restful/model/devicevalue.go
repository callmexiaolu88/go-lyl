package model

//ConditionAndResult 查询条件和结果
type DeviceValueCondition struct {
	Building string
	Floor    string
}

//ConditionAndResult 查询条件和结果
type DeviceFireValueResult struct {
	DeviceID     string
	DeviceLabel  string
	DeviceType   string
	DeviceStatus []string
	MapLocation  *LocationInfo
}

//ConditionAndResult 查询条件和结果
type DeviceWaterValueResult struct {
	DeviceID    string
	DeviceLabel string
	DeviceType  string
	Value       int32
	MapLocation *LocationInfo
}

//LocationInfo 位置信息
type LocationInfo struct {
	S1 []float32
	F1 []float32
	V1 []float32
}

//ConditionAndResult 查询条件和结果
type DeviceFireValueConditionAndResult struct {
	//Search 查询条件
	Search DeviceValueCondition `json:"search,omitempty"`
	//Values 查询结果
	Values []DeviceFireValueResult `json:"values,omitempty"`
}

//ResponseInfo 查询返回信息struct
type DeviceFireValueResponseInfo struct {
	//Code 状态码
	Code int `json:"code,omitempty"`
	//Msg 查询结果的消息
	Msg string `json:"msg,omitempty"`
	//查询结果集合
	Data []DeviceFireValueConditionAndResult `json:"data,omitempty"`
}

//ConditionAndResult 查询条件和结果
type DeviceWaterValueConditionAndResult struct {
	//Search 查询条件
	Search DeviceValueCondition `json:"search,omitempty"`
	//Values 查询结果
	Values []DeviceWaterValueResult `json:"values,omitempty"`
}

//ResponseInfo 查询返回信息struct
type DeviceWaterValueResponseInfo struct {
	//Code 状态码
	Code int `json:"code,omitempty"`
	//Msg 查询结果的消息
	Msg string `json:"msg,omitempty"`
	//查询结果集合
	Data []DeviceWaterValueConditionAndResult `json:"data,omitempty"`
}
