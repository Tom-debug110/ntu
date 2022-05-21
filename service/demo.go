package service

import (
	"time"
)

type demo struct{}

// NewDemo 单例模式创建 demo 实例
func NewDemo() *demo {
	return &demo{}
}

type Daily struct {
	Status
	SignInAt  int64 `json:"sign_in"`
	SignOutAt int64 `json:"sign_out"`
}

type Status struct {
	Code    int    `json:"status_code"`
	Message string `json:"status_msg"`
}

// StatusOk 用户打卡状态-请求成功
func (*demo) StatusOk() Daily {
	return Daily{
		Status: Status{
			Code:    0,
			Message: "success",
		},
		SignInAt:  time.Now().Unix() - 90000,
		SignOutAt: time.Now().Unix() - 3000,
	}
}

// SignOK  签到或签退成功
func (*demo) SignOK() Status {
	return Status{
		Code:    0,
		Message: "success",
	}
}

// SignFail 签到或签退失败
func (*demo) SignFail() Status {
	return Status{
		Code:    1,
		Message: "打卡失败：失败原因",
	}
}

type Total struct {
	AverageHour float64 `json:"average_hour"`
	LateCount   int64   `json:"late_count"`
	LeaveCount  int64   `json:"leave_count"`
}
type Detail struct {
	SignIn  int64 `json:"sign_in"`
	SignOut int64 `json:"sign_out"`
	Date    int64 `json:"date"`
}
type Statistics struct {
	Status
	Total   Total    `json:"total"`
	Details []Detail `json:"details"`
}

func (*demo) Statistics() Statistics {
	return Statistics{
		Status: Status{
			Code:    0,
			Message: "success",
		},
		Total: Total{
			AverageHour: 15.9,
			LateCount:   12,
			LeaveCount:  16,
		},
		Details: []Detail{
			{SignIn: 1653017724, SignOut: 1653019724, Date: time.Now().Unix() - 10000},
			{SignIn: 1653017724, SignOut: 1653019724, Date: time.Now().Unix() - 20000},
			{SignIn: 1653017724, SignOut: 1653019724, Date: time.Now().Unix() - 30000},
			{SignIn: 1653017724, SignOut: 1653019724, Date: time.Now().Unix() - 40000},
			{SignIn: 1653017724, SignOut: 1653019724, Date: time.Now().Unix() - 50000},
			{SignIn: 1653017724, SignOut: 1653019724, Date: time.Now().Unix() - 60000},
			{SignIn: 1653017724, SignOut: 1653019724, Date: time.Now().Unix() - 70000},
			{SignIn: 1653017724, SignOut: 1653019724, Date: time.Now().Unix() - 80000},
		},
	}
}

type User struct {
	ID   int64  `json:"user_id"`
	Name string `json:"name"`
}
type MemberList struct {
	Status
	Users []User `json:"user"`
}

func (*demo) MemberList() MemberList {
	return MemberList{
		Status: Status{
			Code:    0,
			Message: "success",
		},
		Users: []User{
			{ID: 3200421039, Name: "Jack"},
			{ID: 3200421078, Name: "王巴丹"},
			{ID: 3200456787, Name: "步耀恋"},
			{ID: 3190421039, Name: "小伙子"},
			{ID: 3200421039, Name: "Jack1"},
			{ID: 3200421078, Name: "王巴丹1"},
			{ID: 3200456787, Name: "步耀恋1"},
			{ID: 3190421039, Name: "小伙子1"},
			{ID: 3200421039, Name: "Jack2"},
			{ID: 3200421078, Name: "王巴丹2"},
			{ID: 3200456787, Name: "步耀恋2"},
			{ID: 3190421039, Name: "小伙子2"},
			{ID: 3200421039, Name: "Jack3"},
			{ID: 3200421078, Name: "王巴丹3"},
			{ID: 3200456787, Name: "步耀恋3"},
			{ID: 3190421039, Name: "小伙子3"},
		},
	}
}

type User1 struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	TotalHour float64 `json:"total_hour"`
}
type Rank struct {
	Status
	Users []User1 `json:"user"`
}

func (*demo) Rank() Rank {
	return Rank{
		Status: Status{
			Code:    0,
			Message: "success",
		},
		Users: []User1{
			{ID: 3200421039, Name: "Jack", TotalHour: 270.45},
			{ID: 3200421078, Name: "王巴丹", TotalHour: 240.45},
			{ID: 3200456787, Name: "步耀恋", TotalHour: 230.45},
			{ID: 3190421039, Name: "小伙子", TotalHour: 220.45},
			{ID: 3200421039, Name: "Jack1", TotalHour: 210.45},
			{ID: 3200421078, Name: "王巴丹1", TotalHour: 200.45},
			{ID: 3200456787, Name: "步耀恋1", TotalHour: 190.45},
			{ID: 3190421039, Name: "小伙子1", TotalHour: 180.45},
			{ID: 3200421039, Name: "Jack2", TotalHour: 170.45},
			{ID: 3200421078, Name: "王巴丹2", TotalHour: 160.45},
			{ID: 3200456787, Name: "步耀恋2", TotalHour: 150.45},
			{ID: 3190421039, Name: "小伙子2", TotalHour: 140.45},
			{ID: 3200421039, Name: "Jack3", TotalHour: 130.45},
			{ID: 3200421078, Name: "王巴丹3", TotalHour: 120.45},
			{ID: 3200456787, Name: "步耀恋3", TotalHour: 110.45},
			{ID: 3190421039, Name: "小伙子3", TotalHour: 100.45},
		},
	}
}

func (*demo) UserStatus() Status {
	return Status{
		Code:    0,
		Message: "success",
	}
}
