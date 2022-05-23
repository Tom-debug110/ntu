package dao

import (
	"errors"
	"ntu/model"
	"sync"

	"gorm.io/gorm"
)

type userDAO struct{}

var (
	userDAOInstance *userDAO
	userDAOOnce     sync.Once
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
func (*userDAO) Exist(conditions map[string]interface{}) (model.User, bool) {
	var u model.User
	err := db.Model(&model.User{}).Select([]string{"user_id", "name"}).Where(conditions).First(&u).Error
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, false
	}

	return u, true
}

// Create 创建新用户，建立 openID、学号、姓名之间的联系
func (*userDAO) Create(u *model.User) error {
	return db.Model(&model.User{}).Create(u).Error
}

// Update 注册用户时，更新其openid
func (*userDAO) Update(u *model.User) error {
	return db.Model(&model.User{}).Where("user_id=?", u.UserID).
		Select("open_id").
		Updates(map[string]interface{}{"open_id": u.OpenID}).Error
}

// QueryUserByOpenID 通过OpenID 获取用户的信息
func (*userDAO) QueryUserByOpenID(openID string) (model.User, error) {
	var u model.User
	err := db.Model(&model.User{}).Select("user_id", "name").Where("open_id = ?", openID).First(&u).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, err
	}

	return u, nil
}

// QueryUsers 查询用户列表
func (*userDAO) QueryUsers(conditions map[string]interface{}) ([]model.User, error) {
	var u []model.User
	err := db.Model(model.User{}).Where(conditions).Select("user_id", "name").Find(&u).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return u, nil
}
