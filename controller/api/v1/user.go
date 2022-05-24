package v1

import (
	"fmt"
	"net/http"
	"ntu/controller/respones"
	"ntu/model"
	"ntu/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Exist 查询用户是否已经注册，即补充完整自己的信息
func Exist(c *gin.Context) {
	openID := c.Request.Header.Get("x-wx-openid")
	resp := service.NewUserService().Exist(openID)
	c.JSON(http.StatusOK, resp)
}

// Register 用户注册，实现openID和userID、姓名的绑定
func Register(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, respones.ParamsInvalid)
		return
	}
	openID := c.Request.Header.Get("x-wx-openid")
	name := c.Query("name")
	if len(name) > 15 || len(name) == 0 {
		c.JSON(http.StatusBadRequest, respones.ParamsInvalid)
		return
	}

	fmt.Println("controller->user->34:", userID, openID, name)
	u := model.User{
		OpenID: openID,
		UserID: userID,
		Name:   name,
	}
	resp := service.NewUserService().Register(&u)
	c.JSON(http.StatusOK, resp)
}

// List 成员列表
func List(c *gin.Context) {
	resp := service.NewUserService().List()
	c.JSON(http.StatusOK, resp)
}
