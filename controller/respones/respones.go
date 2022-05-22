package respones

type Status struct {
	Code    int
	Message string
}

var OK = Status{Code: 0, Message: "success"}
var ParamsInvalid = Status{Code: -1, Message: "请求参数错误"}
