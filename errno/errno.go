package errno

type Errno struct {
	Code    int
	Message string
}

var (
	//用户相关 101开头
	ErrUserNotExist      = &Errno{Code: 10101, Message: "非实验室成员"}
	ErrUserRegisterFail  = &Errno{Code: 10102, Message: "用户注册失败"}
	ErrQueryUserInfoFail = &Errno{Code: 10103, Message: "查询用户信息失败"}
	ErrQueryUserListFail = &Errno{Code: 10104, Message: "查询用户列表失败"}

	//打卡记录相关 102开头
	ErrRecordQueryFail = &Errno{Code: 10201, Message: "查询记录失败"}
	ErrSignInFail      = &Errno{Code: 10202, Message: "签到失败"}
	ErrSignOutFail     = &Errno{Code: 10203, Message: "签退失败"}
)
