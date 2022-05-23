package main

import (
	"ntu/config"
	"ntu/dao"
	"ntu/router"
)

func main() {
	dao.InitDB()
	r := router.InitRouter()
	r.Run(config.Port)
}
