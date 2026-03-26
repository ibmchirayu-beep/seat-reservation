package main

import "github.com/gin-gonic/gin"

func getSeats(c *gin.Context) {
	c.JSON(200, seats)
}