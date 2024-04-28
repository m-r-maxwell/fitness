package main

import (
	"fitness/models"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

func main() {
	// find out what day it is
	cd := time.Now().Weekday()
	if cd == time.Tuesday || cd == time.Thursday || cd == time.Wednesday || cd == time.Saturday {
		wfd := make([]models.Exercise, 4)
		wfds := make([]string, 4)
		// we need to pick 4 random exercises from the list
		seed := time.Now().UnixNano()
		rdr := rand.New(rand.NewSource(seed))
		listLen := len(models.List)
		wfd[0] = models.List[rdr.Intn(listLen)]
		wfd[1] = models.List[rdr.Intn(listLen)]
		wfd[2] = models.List[rdr.Intn(listLen)]
		wfd[3] = models.List[rdr.Intn(listLen)]
		for x, v := range wfd {
			modText := strings.Join(v.Modifications, ", ")
			wStr := fmt.Sprintf("# of Reps: %d\n", v.Reps)
			if v.WorkingTime > 0 {
				wStr = fmt.Sprintf("Working Time: %d seconds each\n", v.WorkingTime)
			}
			str := fmt.Sprintf("Exercise %d: %s \nSets: %d \n", x+1, v.Name, v.Sets)
			str += wStr
			str += fmt.Sprintf("Description: %s\nPossible Modificatiosn: %v", v.Description, modText)
			wfds[x] = str
		}
		// print out the exercises
		final := "Exercise List\n------------------------------------------------\n"
		for _, v := range wfds {
			final += fmt.Sprintf("%s\n", v)
			final += "------------------------------------------------\n"
		}

		// set up the email
		sender := "Your-email-here"
		password := "Your-password-here"
		recipient := "Your-email-here"

		// send the email
		m := gomail.NewMessage()
		m.SetHeader("From", sender)
		m.SetHeader("To", recipient)
		m.SetHeader("Subject", "Workouts For Today")
		m.SetBody("text/plain", final)
		d := gomail.NewDialer("smtp.goes.here.com", 1000, sender, password)
		if err := d.DialAndSend(m); err != nil {
			log.Fatal(err)
		}

	}

}
