package third_party

import (
	"github.com/subosito/twilio"
	"log"
	"strconv"
)

type TwilioClient struct {
	Twilio *twilio.Twilio
}

func NewTwilioCLient() *TwilioClient {
	t := twilio.NewTwilio("AC8ae5976ecd3c5b1d8d6ee9aea8cdab05", "3b5713b9e11455007ecf00d3064d8bc1")
	return &TwilioClient{Twilio:t}
}

func (t *TwilioClient) SendSMS(phone int, otp int) error {
	smsParams := twilio.SMSParams{}
	resp, err := t.Twilio.SendSMS("+12057514403", "+91" + strconv.Itoa(phone), "OTP " + strconv.Itoa(otp),smsParams)
	log.Println("Response >>>", resp)
	if err != nil {
		return err
	}
	return nil
}