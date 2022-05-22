package dao

import (
	"ntu/model"
	"sync"
)

type userDAO struct{}

var (
	userDAOInstance *userDAO
	userDAOOnce     *sync.Once
)

// NewUserDAOInstance 单例模式创建 userDAO 实例
func NewUserDAOInstance() *userDAO {
	userDAOOnce.Do(
		func() {
			userDAOInstance = &userDAO{}
		},
	)
	return userDAOInstance
}

// Exist 查询满足条件的用户是否存在
// 存在返回true 不存在返回false
// 如果出现错误，返回 false
func (*userDAO) Exist(conditions map[string]interface{}) bool {
	var count int64
	err := db.Model(&model.User{}).Select("id").Where(conditions).Count(&count).Error
	if err != nil {
		return false
	}

	return count > 0
}

// Create 创建新用户，建立 openID、学号、姓名之间的联系
func (*userDAO) Create(u *model.User) error {
	return db.Model(&model.User{}).Create(u).Error
}
