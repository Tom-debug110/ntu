package service

import (
	"fmt"
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
func NewUserService() *userService {
	userServiceOnce.Do(
		func() {
			userServiceInstance = &userService{}
		})
	return userServiceInstance
}

// Exist 用户是否存在
// 判断数据库中userID和OpenID绑定
func (*userService) Exist(openID string) respones.Exist {
	user, ok := dao.NewUserDAOInstance().Exist(map[string]interface{}{
		"open_id": openID,
	})

	if !ok {
		return respones.Exist{
			Status: respones.Status{Code: errno.ErrUserNotExist.Code, Message: errno.ErrUserNotExist.Message},
		}
	}

	return respones.Exist{Status: respones.OK, User: user}
}

func (*userService) Register(u *model.User) respones.Status {

	fmt.Println("service->user->57:", u)
	err := dao.NewUserDAOInstance().Update(u)
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
