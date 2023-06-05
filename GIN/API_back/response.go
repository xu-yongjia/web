package api2

// DataList 带有总数的Data结构
// TokenData 带有token的Data结构
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

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
