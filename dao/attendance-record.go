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
func (*attendanceDao) QueryRecords(userID int64, expr string) ([]model.AttendanceRecord, error) {
	var a []model.AttendanceRecord
	err := db.Model(&model.AttendanceRecord{}).
		Where(map[string]interface{}{"user_id": userID}, gorm.Expr(expr)).
		Omit("id").
		Find(&a).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return a, nil
}

func (*attendanceDao) QuerySingleRecord(userID int64, expr string) (model.AttendanceRecord, error) {
	var a model.AttendanceRecord
	err := db.Model(&model.AttendanceRecord{}).
		Where(map[string]interface{}{"user_id": userID}, gorm.Expr(expr)).
		Omit("id").
		First(&a).Error

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
func (*attendanceDao) Update(expr string, field map[string]interface{}) error {
	return db.Model(&model.AttendanceRecord{}).Where(gorm.Expr(expr)).Updates(field).Error
}

// QueryTotalHour 查询一定条件下用户的总工时
func (*attendanceDao) QueryTotalHour(userID int64, expr string) (float64, error) {
	var r float64
	var c int64
	resultDB := db.Model(&model.AttendanceRecord{}).
		Where(map[string]interface{}{"user_id": userID}, gorm.Expr(expr)).
		Select("sum(timeStampDiff(minute,sign_in_at,sign_out_at))").
		Count(&c)
	if resultDB.Error != nil && errors.Is(resultDB.Error, gorm.ErrRecordNotFound) || c == 0 {
		return 0, resultDB.Error
	}
	resultDB.First(&r)
	return r, nil
}

// AverageStatistics 平均时长
// 精确到分钟
func (*attendanceDao) AverageStatistics(userID int64, expr string) (float64, error) {
	var result float64
	err := db.Model(&model.AttendanceRecord{}).
		Where(model.AttendanceRecord{UserID: userID}, gorm.Expr(expr)).
		Select("AVG(timeStampDiff(minute,sign_in_at,sign_out_at))").
		First(&result).Error
	if err != nil {
		return 0, err
	}
	return result, nil
}

// LateCountStatistics 迟到次数
func (*attendanceDao) LateCountStatistics(userID int64, expr string) (int64, error) {
	var count int64
	err := db.Model(&model.AttendanceRecord{}).
		Where(map[string]interface{}{"user_id": userID}, gorm.Expr(expr)).
		Select("count(*)").
		First(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// LeaveCountStatistics 早退次数
func (*attendanceDao) LeaveCountStatistics(userID int64, expr string) (int64, error) {
	var count int64
	err := db.Model(&model.AttendanceRecord{}).Where(map[string]interface{}{"user_id": userID},
		gorm.Expr(expr)).
		Select("count(*)").
		First(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return count, nil
}
