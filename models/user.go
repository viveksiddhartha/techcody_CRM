package models

import (
	"SV_CRM/common/utility"
	"fmt"
	"time"
)

type User struct {
	UUID              string `json:"uuid" bson:"uuid"`
	Username          string `json:"username" bson:"username"`
	FirstName         string `json:"firstName" bson:"firstName"`
	LastName          string `json:"lastName" bson:"lastName"`
	Email             string `json:"email" bson:"email"`
	PasswordHash      string `json:"passwordHash" bson:"passwordHash"`
	TimestampCreated  int64  `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64  `json:"timestampModißfied" bson:"timestampModified"`
}

func NewUser(username string, firstName string, lastName string, email string, password string) *User {

	fmt.Printf("This is NewUser arg parameter print %v %v %v %v", username, firstName, lastName, email, password)
	passwordHash := utility.SHA256OfString(password)
	now := time.Now()
	unixTimestamp := now.Unix()
	u := User{UUID: utility.GenerateUUID(), Username: username, FirstName: firstName, LastName: lastName, Email: email, PasswordHash: passwordHash, TimestampCreated: unixTimestamp}
	fmt.Printf("this is u %v", u)
	return &u
}
