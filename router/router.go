package router

import (
	"ntu/controller/demo"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/attendance-record/status", demo.Status)
	r.POST("/attendance-record/sign-in", demo.SignOk)
	r.GET("/attendance-record/statistics", demo.Statistics)
	r.GET("/attendance-record/member-list", demo.MemberList)
	r.GET("/attendance-record/rank", demo.Rank)
	r.GET("/attendance-record/user-status", demo.UserStatus)

	return r
}
