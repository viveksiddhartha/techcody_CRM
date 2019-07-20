package models

import (
	"SV_CRM/common/utility"
	"fmt"
	"time"
)

type Profile struct {
	UUID              string `json:"uuid" bson:"uuid"`
	CoEntity          string `json:"CoEntity" bson:"CoEntity"`
	Profilename       string `json:"Profilename" bson:"Profilename"`
	FirstName         string `json:"firstName" bson:"firstName"`
	LastName          string `json:"lastName" bson:"lastName"`
	Email             string `json:"email" bson:"email"`
	ContactNo         string `json:"ContactNo" bson:"ContactNo"`
	PasswordHash      string `json:"passwordHash" bson:"passwordHash"`
	TimestampCreated  int64  `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64  `json:"timestampModißfied" bson:"timestampModified"`
}

func NewProfile(password string) *Profile {

	//func NewProfile(w http.ResponseWriter, r *http.Request) *Profile {
	u := Profile{}
	//fmt.Printf("This is NewProfile arg parameter print %v %v %v %v", Profilename, firstName, lastName, email, password)
	passwordHash := utility.SHA256OfString(u.PasswordHash)
	now := time.Now()
	unixTimestamp := now.Unix()
	b := Profile{CoEntity: u.CoEntity, Profilename: u.Profilename, FirstName: u.FirstName, LastName: u.LastName, Email: u.Email, ContactNo: u.ContactNo, PasswordHash: passwordHash, TimestampCreated: unixTimestamp}
	fmt.Printf("this is u %v", b)
	return &b
}
