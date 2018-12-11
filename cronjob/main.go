package main

import (
	"fmt"
	"github.com/sfreiberg/gotwilio"
)

func main() {
	accountSid := "ACf44edc807510af45d5da7ce76d34f373"
	authToken := "5fda003f65eeec1f1238c298c18cc082"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+15853674765"
	to := "+917982830291"
	message := "this is cron job test"
	resp, exp, err := twilio.SendSMS(from, to, message, "", "")
	fmt.Printf("resp:%v \n exp:%v \n err:%v", resp, exp, err)
}
