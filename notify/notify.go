package notify

import (
	"fmt"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func Notify(msg string) int {
	accountSid := "AC65ddca20c273229aa7d74d489876da6d"
	authToken := "d62083932546d143e04e9f09c0a2dfea"
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("+17324477247")
	params.SetFrom("+19707143698")
	params.SetBody(msg)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	} else {
		fmt.Println("SMS sent successfully!")
		return 200
	}
}
