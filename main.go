package main

import (
	"ntu/config"
	"ntu/dao"
	"ntu/router"
)

func main() {
	dao.InitDB()
	r := router.InitRouter()
	err := r.Run(config.Port)
	if err != nil {
		panic(err)
	}

}
