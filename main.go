package main

import (
	"fmt"
	"net/http"

)

func scheduleAlarm(response http.ResponseWriter, req *http.Request) {
	fmt.Println("Server")
}

func main() {
	fmt.Println("Starting scheduling server ----->")
	http.HandleFunc("/alarm", scheduleAlarm)
	http.ListenAndServe(":10000",nil)
}
