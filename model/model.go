package model

import "time"

// User 用户实体
type User struct {
	ID     uint
	OpenID string `gorm:"type:char(30);unique;not null;comment:小程序openid"`
	UserID int64  `gorm:"unique;not null;comment:用户学号或者工号"`
	Name   string `gorm:"type:char(10);not null;comment:用户姓名"`
}

// AttendanceRecord 打卡记录
type AttendanceRecord struct {
	ID        uint64
	UserID    int64     `gorm:"not null;comment:打卡的用户id"`
	Date      time.Time `gorm:"not null;comment:计划打卡时间"`
	SignInAt  time.Time `gorm:"comment:签到时间"`
	SignOutAt time.Time `gorm:"comment:签退时间"`
}
