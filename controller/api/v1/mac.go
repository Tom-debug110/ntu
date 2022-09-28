package v1

import (
	"net/http"
	"ntu/controller/respones"
	"ntu/service"

	"github.com/gin-gonic/gin"
)

func MacAddressUpdate(c *gin.Context) {
	mac := c.Query("mac_address")
	err := service.NewMac().Update(mac)
	if err != nil {
		c.JSON(http.StatusOK, respones.Status{
			Code:    -1,
			Message: "mac 地址更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, respones.OK)
}

func MacAddressQuery(c *gin.Context) {
	mac, err := service.NewMac().Query()
	if err != nil {
		c.JSON(http.StatusOK, respones.Status{
			Code:    4409,
			Message: "mac 地址查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      0,
		"mac_address": mac,
	})
}
