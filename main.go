package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Hello {
	Name string
}

func scheduleAlarm(c *gin.Context) {
	fmt.Println("Getting request")
	obj := Hello()
	obj.Name = "arijit"
	c.JSON(200,obj)
}

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")
	fmt.Println("Starting scheduling server -----> ", port)

	router.GET("/alarm", scheduleAlarm)
	router.Run(":" + port)
}
