package templatemethod

import "fmt"

func Main() {
	// var otp *Otp
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	// otp := smsOTP.(*Otp)
	// not possible

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)

}
