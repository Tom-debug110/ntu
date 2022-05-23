package router

import (
	v1 "ntu/controller/api/v1"
	"ntu/controller/demo"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/attendance-record/user/status/", v1.Exist)  //用户状态
	r.POST("/attendance-record/register/", v1.Register) //用户注册、完善信息
	r.GET("/attendance-record/status/", v1.Status)      //打卡状态
	r.POST("/attendance-record/sign-in/", v1.SignIn)    //签到
	r.POST("/attendance-record/sign-out/", v1.SignOut)  //签退
	r.GET("/test/", v1.TestLate)                        //测试

	r.GET("/attendance-record/statistics/", demo.Statistics)
	r.GET("/attendance-record/member-list/", demo.MemberList)
	r.GET("/attendance-record/rank/", demo.Rank)

	return r
}
