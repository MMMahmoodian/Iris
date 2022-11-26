package controllers

import (
	"github.com/MMMahmoodian/alarm/conf/rabbit_mq"
	"github.com/MMMahmoodian/alarm/models"
	"github.com/MMMahmoodian/alarm/services/alarm"
	"github.com/MMMahmoodian/alarm/services/queue"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TelegramController struct{}

var dispatcher = alarm.Dispatcher{}
var publisher = queue.RMQPublisher{}

func (TelegramController) Dispatch(c *gin.Context) {
	message := new(models.TelegramMessage)
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": "",
			"errors":   err.Error(),
		})
		return
	}
	if err := dispatcher.Dispatch(publisher, message, rabbit_mq.TelegramQueue); err != nil {
		c.JSON(500, gin.H{
			"messages": "",
			"errors":   err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"messages": "success",
		"errors":   "",
	})
}

func (TelegramController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "pong",
		"errors":   "",
	})
}
