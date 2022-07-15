package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string
}

func AddJob(c *gin.Context) {
	var response Response
	response.Message = "SUCCESS"
	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")
	fmt.Println("Starting scheduling server -----> ", port)
	router.LoadHTMLGlob("templates/alarm.tmpl")
	router.GET("/alarm", AddJob)
	router.Run(":" + port)
}
