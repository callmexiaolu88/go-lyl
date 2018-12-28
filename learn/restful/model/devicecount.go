package model

//ConditionAndResult 查询条件和结果
type DeviceCountCondition struct {
	Building string
	Floor    string
}

//ConditionAndResult 查询条件和结果
type DeviceCountResult struct {
	Name  string
	Value int
}

//ConditionAndResult 查询条件和结果
type DeviceCountConditionAndResult struct {
	//Search 查询条件
	Search DeviceCountCondition `json:"search,omitempty"`
	//Values 查询结果
	Values []DeviceCountResult `json:"values,omitempty"`
}

//ResponseInfo 查询返回信息struct
type DeviceCountResponseInfo struct {
	//Code 状态码
	Code int `json:"code,omitempty"`
	//Msg 查询结果的消息
	Msg string `json:"msg,omitempty"`
	//查询结果集合
	Data []DeviceCountConditionAndResult `json:"data,omitempty"`
}
