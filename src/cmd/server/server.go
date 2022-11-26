package main

import (
	"github.com/MMMahmoodian/alarm/conf"
	"github.com/MMMahmoodian/alarm/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	conf.Initialize()
}

func main() {
	r := gin.Default()
	telegramController := controllers.TelegramController{}
	
	r.GET("/telegram/ping", telegramController.Ping)
	r.POST("/telegram/", telegramController.Dispatch)

	r.Run()
}
