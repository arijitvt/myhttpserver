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
	http.ListenAndServer(":10000",nil)
}
