package model

//ConditionAndResult 查询条件和结果
type DeviceDetailCondition struct {
	Building string
	Floor    string
}

//ConditionAndResult 查询条件和结果
type DeviceDetailResult struct {
	Max        int32
	Min        int32
	Unit       string
	Value      int32
	DeviceType string
}

//ConditionAndResult 查询条件和结果
type DeviceDetailConditionAndResult struct {
	//Search 查询条件
	Search DeviceDetailCondition `json:"search,omitempty"`
	//Values 查询结果
	Values []DeviceDetailResult `json:"values,omitempty"`
}

//ResponseInfo 查询返回信息struct
type DeviceDetailResponseInfo struct {
	//Code 状态码
	Code int `json:"code,omitempty"`
	//Msg 查询结果的消息
	Msg string `json:"msg,omitempty"`
	//查询结果集合
	Data []DeviceDetailConditionAndResult `json:"data,omitempty"`
}
