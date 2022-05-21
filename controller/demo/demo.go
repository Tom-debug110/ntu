package demo

import (
	"net/http"
	"ntu/service"

	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, service.NewDemo().StatusOk())
}

func SignOk(c *gin.Context) {
	c.JSON(http.StatusOK, service.NewDemo().SignOK())
}

func Statistics(c *gin.Context) {
	c.JSON(http.StatusOK, service.NewDemo().Statistics())
}

func MemberList(c *gin.Context) {
	c.JSON(http.StatusOK, service.NewDemo().MemberList())
}

func Rank(c *gin.Context) {
	c.JSON(http.StatusOK, service.NewDemo().Rank())
}

func UserStatus(c *gin.Context) {
	c.JSON(http.StatusOK, service.NewDemo().UserStatus())
}
