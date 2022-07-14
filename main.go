package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func scheduleAlarm(response http.ResponseWriter, req *http.Request) {
	fmt.Println("Server")
	output := "Hello From the server"
	fmt.Fprintf(response, "Hello %q ", output)
}

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")
	fmt.Println("Starting scheduling server -----> ", port)
	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	httpServer.ListenAndServe()
}
