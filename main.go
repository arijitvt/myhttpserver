package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arijitvt/myhttpserver/notify"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string
}

func AddJob(c *gin.Context) {
	var response Response
	response.Message = "SUCCESS"
	result := notify.Notify(response.Message)
	if result == 200 {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusBadRequest, "Could not send request")
	}

}

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")
	fmt.Println("Starting scheduling server -----> ", port)
	router.LoadHTMLGlob("templates/alarm.tmpl")
	router.GET("/alarm", AddJob)
	router.Run(":" + port)
}
