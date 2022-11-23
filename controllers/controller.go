package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Dispatch(c *gin.Context)
	Ping(c *gin.Context)
}
