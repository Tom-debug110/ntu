package v1

import (
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
	date, err := timeStampToTime(c.Query("date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, respones.ParamsInvalid)
		return
	}
	res := service.NewRecordService().Status(date)
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

func TestLate(c *gin.Context) {
	service.NewRecordService().LateCount(3200421039)
}
