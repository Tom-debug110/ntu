package service

import (
	"ntu/controller/respones"
	"ntu/dao"
	"ntu/errno"
	"ntu/model"
	"sync"
)

type userService struct {
	UserID int64
	OpenID string
}

var (
	userServiceInstance *userService
	userServiceOnce     sync.Once
)

// NewUserService 创建 userService 实例
func NewUserService(userID int64, openID string) *userService {
	userServiceOnce.Do(
		func() {
			userServiceInstance = &userService{
				UserID: userID,
				OpenID: openID,
			}
		})
	return userServiceInstance
}

// Exist 用户是否存在
// 判断数据库中userID和OpenID绑定
func (u *userService) Exist() respones.Status {
	res := dao.NewUserDAOInstance().Exist(map[string]interface{}{
		"open_id": u.OpenID,
	})

	if res {
		return respones.OK
	}

	return respones.Status{Code: errno.ErrUserNotExist.Code, Message: errno.ErrUserNotExist.Message}
}

func (u *userService) Register(name string) respones.Status {
	user := model.User{
		OpenID: u.OpenID,
		UserID: u.UserID,
		Name:   name,
	}
	err := dao.NewUserDAOInstance().Create(&user)
	if err != nil {
		return respones.Status{Code: errno.ErrUserRegisterFail.Code, Message: errno.ErrUserRegisterFail.Message}
	}
	return respones.OK
}

func (*userService) List() respones.UserList {
	u, err := dao.NewUserDAOInstance().QueryUsers(map[string]interface{}{})
	if err != nil {
		return respones.UserList{Status: handleErr(errno.ErrQueryUserListFail)}
	}

	return respones.UserList{Status: respones.OK, Users: u}
}
