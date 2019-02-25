package Tools

import (
	"encoding/json"
)

// 返回json
type Msg struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data []string `json:"data"`
}

func ResJson(code int,msg string,data []string)[]byte{
	vc := Msg{code,msg,data}
	res,_ := json.Marshal(vc)
	return res
}