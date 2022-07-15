package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/arijitvt/myhttpserver/notify"
	"github.com/gin-gonic/gin"
)

func AddJob(c *gin.Context) {
	msg := c.Query("message")
	duration, _ := time.ParseDuration(c.Query("duration"))
	go func() {
		notify.Notify(msg, duration)
	}()
	c.JSON(http.StatusOK, "Scheduled")
}

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")
	fmt.Println("Starting scheduling server -----> ", port)
	// router.LoadHTMLGlob("templates/alarm.tmpl")
	router.GET("/alarm", AddJob)
	router.Run(":" + port)
}
