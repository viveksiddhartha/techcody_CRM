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

func NewUser(password string) *User {

	//func NewUser(w http.ResponseWriter, r *http.Request) *User {
	u := User{}
	//fmt.Printf("This is NewUser arg parameter print %v %v %v %v", username, firstName, lastName, email, password)
	passwordHash := utility.SHA256OfString(u.PasswordHash)
	now := time.Now()
	unixTimestamp := now.Unix()
	b := User{Username: u.Username, FirstName: u.FirstName, LastName: u.LastName, Email: u.Email, PasswordHash: passwordHash, TimestampCreated: unixTimestamp}
	fmt.Printf("this is u %v", b)
	return &b
}
