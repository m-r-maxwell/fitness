package main

import (
	h "fitness/helpers"
	"time"
)

func main() {
	// set up the email
	sender := "your-email-here"
	password := "your-password-here"

	// find out what day it is
	cd := time.Now().Weekday()

	// create the email body
	final := h.CreateEmailBody()

	switch cd {
	case time.Monday:
		// add func here to send email
		h.SendEmail(sender, password, final)
	case time.Tuesday:
		// add func here to send email
		// h.SendEmail(sender, password, final)
	case time.Wednesday:
		// add func here to send email
		// h.SendEmail(sender, password, final)
	case time.Thursday:
		// add func here to send email
		// h.SendEmail(sender, password, final)
	case time.Friday:
		// add func here to send email
		// h.SendEmail(sender, password, final)
	case time.Saturday:
		// add func here to send email
		// h.SendEmail(sender, password, final)
	case time.Sunday:
		// add func here to send email
		// h.SendEmail(sender, password, final)
	default:
		// skip
	}

}
