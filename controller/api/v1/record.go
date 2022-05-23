package v1

import (
	"fmt"
	"net/http"
	"ntu/controller/respones"
	"ntu/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 时间戳转time.Time 类型
func timeStampToTime(timeStamp string) (time.Time, error) {
	date, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		return time.Now(), err
	}

	return time.Unix(date, 0), nil
}

// Status 用户打卡状态
func Status(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, respones.ParamsInvalid)
		return
	}
	res := service.NewRecordService().Status(userID)
	fmt.Println(res)
	c.JSON(http.StatusOK, res)
}

// SignIn 签到接口
func SignIn(c *gin.Context) {
	openid := c.Request.Header.Get("x-wx-openid")
	resp := service.NewRecordService().SignIn(openid)
	c.JSON(http.StatusOK, resp)
}

func SignOut(c *gin.Context) {
	openid := c.Request.Header.Get("x-wx-openid")
	resp := service.NewRecordService().SignOut(openid)
	c.JSON(http.StatusOK, resp)
}

// Statistics 打卡统计
func Statistics(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, respones.ParamsInvalid)
		return
	}

	resp := service.NewRecordService().Statistics(userID)
	c.JSON(http.StatusOK, resp)
}

// Rank 排行榜
func Rank(c *gin.Context) {
	resp := service.NewRecordService().Rank()
	c.JSON(http.StatusOK, resp)
}
