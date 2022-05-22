package errno

type Errno struct {
	Code    int
	Message string
}

var (
	//用户相关 101开头
	ErrUserNotExist     = &Errno{Code: 10101, Message: "非实验室成员"}
	ErrUserRegisterFail = &Errno{Code: 10102, Message: "用户注册失败"}
)
