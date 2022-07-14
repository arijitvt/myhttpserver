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
	http.HandleFunc("/alarm", scheduleAlarm)
	port := os.Getenv("PORT")
	fmt.Println("Starting scheduling server -----> ", port)
	http.ListenAndServe(":"+port,nil)
}
