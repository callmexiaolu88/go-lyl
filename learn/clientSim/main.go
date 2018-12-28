package main

import fmt "fmt"

func main() {
	StartWaterCalling()
	fmt.Println("-------------------------")
	StartFireCalling()
}

// func requestData(bytes []byte, path string, searchs, values *[]interface{}) (pb.OperateResult, string, error) {
// 	body, err := json.Marshal(QueryInfo{
// 		Method: "POST",
// 		Path:   path,
// 		Data:   bytes,
// 	})
// 	if err != nil {
// 		return pb.OperateResult_Fault, "", err
// 	}
// 	rsp, err := protocolhandler.QueryDevice(body)
// 	if err != nil {
// 		return pb.OperateResult_Fault, "", err
// 	}

// 	rspinfo := &ResponseInfo{}
// 	err = json.Unmarshal(rsp, rspinfo)
// 	if err != nil {
// 		return pb.OperateResult_Fault, "", err
// 	}

// 	if len(searchs) == len(values) == len(rspinfo.Data) {
// 		for k, dt := range rspinfo.Data {
// 			err = json.Unmarshal(dt.Search, (*searchs)[k])
// 			if err != nil {
// 				return err
// 			}
// 			err = json.Unmarshal(dt.Values, (*values)[k])
// 			if err != nil {
// 				return pb.OperateResult_Fault, "", err
// 			}
// 		}
// 		return codeToResult(rspinfo.Code), rspinfo.Msg, err
// 	}
// 	return pb.OperateResult_Fault, "", errors.New("count of queryed and reply count is not match")
// }
