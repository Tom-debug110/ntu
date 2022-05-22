package dao

import (
	"ntu/model"
	"sync"
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

func (*attendanceDao) Query(conditions map[string]interface{}, field ...string) []model.AttendanceRecord {

}
