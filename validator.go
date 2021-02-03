package utils

import (
	"log"
	"regexp"
)

func EmailValidator(email string) bool {
	log.Println("email verify: ", email)
	var re = regexp.MustCompile(ReEmail)
	return len(re.FindStringIndex(email))>0
}
