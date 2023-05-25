package api1

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// 用于产生返回出错信息
func ERRRESPONSE(reason string, errcode int) Response {
	return Response{
		Status: errcode,
		Msg:    reason,
	}
}

// 用于返回成功信息以及附带的数据
func SUCCESSRESPONSE(data interface{}) Response {
	return Response{
		Msg:    "OK",
		Status: 200,
		Data:   data,
	}
}

func SUCCESSRESPONSE_NODATA() Response {
	return Response{
		Msg:    "OK",
		Status: 200,
	}
}

// Response 基础序列化器
type EResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
