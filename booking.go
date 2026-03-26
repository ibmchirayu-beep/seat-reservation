package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func bookSeat(c *gin.Context) {
	seatID, _ := strconv.Atoi(c.Query("seat_id"))

	mutex.Lock()
	defer mutex.Unlock()

	for i := range seats {
		if seats[i].ID == seatID {
			if seats[i].Booked {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Seat already booked"})
				return
			}
			seats[i].Booked = true
			c.JSON(200, gin.H{"message": "Seat booked"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Seat not found"})
}