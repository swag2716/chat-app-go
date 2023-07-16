package otpControllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func sendOTPHandler(c *gin.Context) {
	phoneNumber := ""

	otp := GenerateOTP(6)

	err := SendOTP(phoneNumber, otp)
	if err != nil {
		log.Fatal("Failed to send OTP:", err)
	}

	log.Println("OTP sent successfully!")
}
