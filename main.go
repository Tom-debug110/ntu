package main

import "ntu/router"

func main() {
	r := router.InitRouter()
	r.Run(":80")
}
