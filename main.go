package main

import (
	"os"
	"sync"
    "github.com/gin-gonic/gin"
)

type Seat struct {
	ID     int  `json:"id"`
	Booked bool `json:"booked"`
}

var seats []Seat
var mutex sync.Mutex

func main() {
	// Create 20 seats
	for i := 1; i <= 20; i++ {
		seats = append(seats, Seat{ID: i, Booked: false})
	}

	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/seats", getSeats)
	router.POST("/book", bookSeat)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}