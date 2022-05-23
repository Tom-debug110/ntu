package model

import "time"

// User 用户实体
type User struct {
	ID     uint   `json:"id,omitempty"`
	OpenID string `gorm:"type:char(30);unique;not null;comment:小程序openid" json:"open_id,omitempty"`
	UserID int64  `gorm:"unique;not null;comment:用户学号或者工号" json:"user_id,omitempty"`
	Name   string `gorm:"type:char(10);not null;comment:用户姓名" json:"name,omitempty"`
}

// AttendanceRecord 打卡记录
type AttendanceRecord struct {
	ID        uint64
	UserID    int64     `gorm:"not null;comment:打卡的用户id"`
	SignInAt  time.Time `gorm:"comment:签到时间"`
	SignOutAt time.Time `gorm:"comment:签退时间"`
}
