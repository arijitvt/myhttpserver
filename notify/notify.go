package notify

import (
	"fmt"
	"time"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func Notify(msg string, wait time.Duration) (int, string) {
	time.Sleep(wait)
	accountSid := "AC65ddca20c273229aa7d74d489876da6d"
	authToken := "a5812398c50015229727b62a7a880c71"
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("+17324477247")
	params.SetFrom("+19707143698")
	params.SetMessagingServiceSid("MG86a3ae4b78a39c4f6a3378a533c90817")
	params.SetBody(msg)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err.Error()
	} else {
		fmt.Println("SMS sent successfully!")
		return 200, "SUCCESS"
	}
}
