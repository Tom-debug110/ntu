package respones

import "ntu/model"

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

type UserList struct {
	Status
	Users []model.User `json:"user"`
}

type UserAPI struct {
	UserID    int64
	Name      string
	TotalHour float64
}
type Rank struct {
	Status
	Users []UserAPI `json:"user"`
}

type Total struct {
	Average    float64 `json:"average_hour"`
	LateCount  int64   `json:"late_count"`
	LeaveCount int64   `json:"leave_count"`
}
type Statistics struct {
	Status
	Total   Total                    `json:"total"`
	Records []model.AttendanceRecord `json:"details"`
}
