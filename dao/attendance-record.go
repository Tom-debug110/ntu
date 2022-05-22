package dao

import (
	"errors"
	"ntu/model"
	"sync"
	"time"

	"gorm.io/gorm"
)

// 单例模式
//var (
//	postDao  *PostDao
//	postOnce sync.Once
//)
//
//func NewPostDaoInstance() *PostDao {
//	postOnce.Do(
//		func() {
//			postDao = &PostDao{}
//		})
//	return postDao
//}

type attendanceDao struct{}

var (
	attendanceDaoInstance *attendanceDao
	attendanceDaoOnce     sync.Once
)

func NewAttendDAOInstance() *attendanceDao {
	attendanceDaoOnce.Do(
		func() {
			attendanceDaoInstance = &attendanceDao{}
		})
	return attendanceDaoInstance
}

// QueryRecords 批量查询一批记录
func (*attendanceDao) QueryRecords(conditions map[string]interface{}, field ...string) ([]model.AttendanceRecord, error) {
	var a []model.AttendanceRecord
	err := db.Model(&model.AttendanceRecord{}).Where(conditions).Select(field).Find(&a).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return a, nil
}

func (*attendanceDao) QuerySingleRecord(conditions map[string]interface{}, field ...string) (model.AttendanceRecord, error) {
	var a model.AttendanceRecord
	err := db.Model(&model.AttendanceRecord{}).Where(conditions).Select(field).First(&a).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.AttendanceRecord{}, err
	}

	return a, nil
}

// Create 创建一条签到记录
// 时间自动创建
func (*attendanceDao) Create(userID int64) error {

	return db.Model(&model.AttendanceRecord{}).Create(map[string]interface{}{
		"user_id":    userID,
		"sign_in_at": time.Now().Format(time.RFC3339),
	}).Error
}

// Update 更新记录
func (*attendanceDao) Update(userID int64, field map[string]interface{}) error {
	return db.Model(&model.AttendanceRecord{}).Where("user_id = ?", userID).Updates(field).Error
}
