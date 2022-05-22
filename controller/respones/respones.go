package respones

type Status struct {
	Code    int    `json:"status_code"`
	Message string `json:"status_msg"`
}

var OK = Status{Code: 0, Message: "success"}
var ParamsInvalid = Status{Code: -1, Message: "请求参数错误"}

type Record struct {
	Status
	SignIn  int64 `json:"sign_in"`
	SingOut int64 `json:"sing_out"`
}
