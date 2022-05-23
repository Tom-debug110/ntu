package service

import (
	"errors"
	"fmt"
	"ntu/controller/respones"
	"ntu/dao"
	"ntu/errno"
	"sort"
	"sync"
	"time"

	"gorm.io/gorm"
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
func (r *recordService) Status(userID int64) respones.Record {
	t := time.Now()
	expr := fmt.Sprintf("year(sign_in_at) = %d AND month(sign_in_at)=%d AND day(sign_in_at)=%d", t.Year(), t.Month(), t.Day())
	res, err := dao.NewAttendDAOInstance().QuerySingleRecord(userID, expr)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return respones.Record{Status: respones.Status{
			Code:    -1,
			Message: err.Error(),
		}}
	}
	if err != nil {
		return respones.Record{Status: respones.Status{
			Code:    errno.ErrRecordQueryFail.Code,
			Message: errno.ErrRecordQueryFail.Message,
		}}
	}
	return respones.Record{Status: respones.OK, SignIn: res.SignInAt.Unix(), SignOut: res.SignOutAt.Unix()}
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
func (*recordService) Rank() respones.Rank {
	t := time.Now()
	expr := fmt.Sprintf("year(sign_in_at) = %d AND month(sign_in_at)=%d AND sign_in_at not null AND sign_out_at not null", t.Year(), t.Month())
	users, err := dao.NewUserDAOInstance().QueryUsers(map[string]interface{}{})
	if err != nil {
		return respones.Rank{Status: handleErr(errno.ErrQueryUserListFail)}
	}
	var u []respones.UserAPI
	for _, i := range users {
		resp, _ := dao.NewAttendDAOInstance().QueryTotalHour(i.UserID, expr)
		u = append(u, respones.UserAPI{
			UserID:    i.UserID,
			Name:      i.Name,
			TotalHour: resp / 60,
		})
	}

	sort.Slice(u, func(i, j int) bool {
		return u[i].TotalHour > u[j].TotalHour
	})
	return respones.Rank{Status: respones.OK, Users: u}
}

// 统计迟到次数
func lateCount(userID int64) int64 {
	t := time.Now()
	expr := fmt.Sprintf("year(sign_in_at) = %d AND month(sign_in_at)=%d AND hour(sign_in_at)>%d", t.Year(), t.Month(), 9)
	res, err := dao.NewAttendDAOInstance().LateCountStatistics(userID, expr)
	if err != nil {
		return 0
	}

	return res
}

// 统计早退次数
func leaveCount(userID int64) int64 {
	t := time.Now()
	expr := fmt.Sprintf("year(sign_in_at) = %d AND month(sign_in_at)=%d AND (sign_out_at IS NULL OR hour(sign_out_at)<21)", t.Year(), t.Month())
	res, err := dao.NewAttendDAOInstance().LeaveCountStatistics(userID, expr)
	if err != nil {
		return 0
	}
	return res
}

// 统计平均打卡时间
func averageStatistics(userID int64) float64 {
	t := time.Now()
	expr := fmt.Sprintf("year(sign_in_at) = %d AND month(sign_in_at)=%d AND sign_out_at IS NOT NULL AND sign_in_at IS NOT NULL", t.Year(), t.Month())
	res, err := dao.NewAttendDAOInstance().AverageStatistics(userID, expr)
	if err != nil {
		return 0
	}
	return res
}

func (*recordService) Statistics(userID int64) respones.Statistics {
	t := time.Now()
	expr := fmt.Sprintf("year(sign_in_at) = %d AND month(sign_in_at)=%d", t.Year(), t.Month())
	average := averageStatistics(userID)
	late := lateCount(userID)
	leave := leaveCount(userID)

	detail, err := dao.NewAttendDAOInstance().QueryRecords(userID, expr)
	if err != nil {
		return respones.Statistics{
			Status: respones.Status{
				Code:    errno.ErrRecordQueryFail.Code,
				Message: errno.ErrRecordQueryFail.Message,
			},
		}
	}

	return respones.Statistics{
		Status: respones.OK,
		Total: respones.Total{
			Average:    average / 60,
			LateCount:  late,
			LeaveCount: leave,
		},
		Records: detail,
	}
}
