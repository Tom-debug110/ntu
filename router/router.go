package router

import (
	v1 "ntu/controller/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/mac-config/update", v1.MacAddressUpdate)      //更新 mac 地址
	r.GET("/mac-config/query", v1.MacAddressQuery)         // 查询mac 地址
	r.GET("/attendance-record/user/status/", v1.Exist)     //用户状态
	r.POST("/attendance-record/register/", v1.Register)    //用户注册、完善信息
	r.GET("/attendance-record/status/", v1.Status)         //打卡状态
	r.POST("/attendance-record/sign-in/", v1.SignIn)       //签到
	r.POST("/attendance-record/sign-out/", v1.SignOut)     //签退
	r.GET("/attendance-record/statistics/", v1.Statistics) //打卡统计
	r.GET("/attendance-record/member-list/", v1.List)      //成员列表
	r.GET("/attendance-record/rank/", v1.Rank)

	return r
}
