package main

import (
	"ntu/dao"
	"ntu/router"
)

func main() {
	dao.InitDB()
	r := router.InitRouter()
	r.Run(":3000")
}
