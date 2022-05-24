package main

import (
	"fmt"
	"ntu/config"
	"ntu/dao"
	"ntu/router"
	"time"
)

func main() {
	dao.InitDB()
	r := router.InitRouter()
	r.Run(config.Port)

	t := time.Now()

	fmt.Println(t.Format("2006-01-02T15:04:05+8:00"))
}
