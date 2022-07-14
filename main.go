package main

import (
	"fmt"
	"net/http"
	"os"
)

func scheduleAlarm(response http.ResponseWriter, req *http.Request) {
	fmt.Println("Server")
	output := "Hello From the server"
	fmt.Fprintf(response,"Hello %q " , output )
}

func main() {
	fmt.Println("Starting scheduling server ----->")
	http.HandleFunc("/alarm", scheduleAlarm)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port,nil)
}
