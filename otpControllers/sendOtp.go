package otpControllers

import (
	"math/rand"
	"time"

	"log"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

const otpDigits = "0123456789"

func GenerateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	otp := make([]byte, length)
	for i := 0; i < length; i++ {
		otp[i] = otpDigits[rand.Intn(len(otpDigits))]
	}
	return string(otp)
}

const accountSID = "YOUR_ACCOUNT_SID"
const authToken = "YOUR_AUTH_TOKEN"
const twilioNumber = "YOUR_TWILIO_PHONE_NUMBER"

func SendOTP(phoneNumber, otp string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})

	messageParams := &twilioApi.CreateMessageParams{}
	messageParams.SetTo(phoneNumber)
	messageParams.SetFrom(twilioNumber)
	messageParams.SetBody("Your OTP is: " + otp)

	_, err := client.Api.CreateMessage(messageParams)
	if err != nil {
		log.Println("Error sending OTP:", err)
		return err
	}

	return nil
}
