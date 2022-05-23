package service

import (
	"fmt"
	"ntu/controller/respones"
	"ntu/dao"
	"ntu/errno"
	"sync"
	"time"
)

type recordService struct{}

var (
	recordServiceInstance *recordService
	recordServiceOnce     sync.Once
)

func handleErr(errType *errno.Errno) respones.Status {
	return respones.Status{
		Code:    errType.Code,
		Message: errType.Message,
	}
}
func NewRecordService() *recordService {
	recordServiceOnce.Do(
		func() {
			recordServiceInstance = &recordService{}
		})

	return recordServiceInstance
}

// Status 当天的打卡状态
func (r *recordService) Status(curTime time.Time) respones.Record {
	res, err := dao.NewAttendDAOInstance().QuerySingleRecord(map[string]interface{}{
		"year(sign_in_at)":  curTime.Year(),
		"month(sign_in_at)": curTime.Month(),
		"day(sign_in_at)":   curTime.Day(),
	}, "sign_in_at", "sign_out_at")
	if err != nil {
		return respones.Record{Status: respones.Status{
			Code:    errno.ErrRecordQueryFail.Code,
			Message: errno.ErrRecordQueryFail.Message,
		}}
	}
	return respones.Record{Status: respones.OK, SignIn: res.SignInAt.Unix(), SingOut: res.SignOutAt.Unix()}
}

// SignIn 签到服务
func (*recordService) SignIn(openID string) respones.Status {
	u, err := dao.NewUserDAOInstance().QueryUserByOpenID(openID)
	if err != nil {
		return handleErr(errno.ErrQueryUserInfoFail)
	}

	err = dao.NewAttendDAOInstance().Create(u.UserID)
	if err != nil {
		return handleErr(errno.ErrSignInFail)
	}

	return respones.OK
}

func (*recordService) SignOut(openID string) respones.Status {
	u, err := dao.NewUserDAOInstance().QueryUserByOpenID(openID)
	if err != nil {
		return handleErr(errno.ErrQueryUserInfoFail)
	}

	err = dao.NewAttendDAOInstance().Update(u.UserID, map[string]interface{}{
		"sign_out_at": time.Now().Format(time.RFC3339),
	})
	if err != nil {
		return handleErr(errno.ErrSignOutFail)
	}

	return respones.OK
}

// Rank 工时排行榜
func (*recordService) Rank(curTime time.Time) respones.Rank {
	users, err := dao.NewUserDAOInstance().QueryUsers(map[string]interface{}{})
	if err != nil {
		return respones.Rank{Status: handleErr(errno.ErrQueryUserListFail)}
	}
	var u []respones.UserAPI
	for _, i := range users {
		resp, _ := dao.NewAttendDAOInstance().QueryTotalHour(map[string]interface{}{
			"user_id":           i.UserID,
			"year(sign_in_at)":  curTime.Year(),
			"month(sign_in_at)": curTime.Month(),
			"sign_out_at":       "not null",
		})
		u = append(u, respones.UserAPI{
			UserID:    i.UserID,
			Name:      i.Name,
			TotalHour: resp,
		})
	}

	return respones.Rank{Status: respones.OK, Users: u}
}

func (*recordService) LateCount(userID int64) int64 {
	t := time.Now()
	expr := fmt.Sprintf("year(sign_in_at) = %d AND month(sign_in_at)=%d AND hour(sign_in_at)>%d", t.Year(), t.Month(), 9)
	res, err := dao.NewAttendDAOInstance().LateCountStatistics(userID, expr)
	if err != nil {
		return 0
	}

	return res
}
