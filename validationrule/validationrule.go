package validationrule

import (
	"log"
	"math/rand"
	"regexp"
	"time"
)

const ProfilenameRegex string = `^@?(\w){1,15}$`
const EmailRegex = `(?i)^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,3})+$`

func CheckProfilenameSyntax(Profilename string) bool {

	validationResult := false
	r, err := regexp.Compile(ProfilenameRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = r.MatchString(Profilename)
	return validationResult
}

func CheckEmailSyntax(email string) bool {
	validationResult := false
	r, err := regexp.Compile(EmailRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = r.MatchString(email)
	return validationResult
}

func GenerateRandomProfilename() string {

	rand.Seed(time.Now().UnixNano())

	ProfilenameLength := rand.Intn(15) + 1

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	b := make([]rune, ProfilenameLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	randomProfilename := string(b)

	zeroOrOne := rand.Intn(2)
	if zeroOrOne == 1 {
		randomProfilename = "@" + randomProfilename
	}
	return randomProfilename
}
