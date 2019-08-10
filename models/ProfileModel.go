package models

type Profile struct {
	UUID              string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	CoEntityID        string `json:"CoEntityID,omitempty" bson:"CoEntityID,omitempty"`
	Profilename       string `json:"Profilename,omitempty" bson:"Profilename,omitempty"`
	FirstName         string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email             string `json:"email,omitempty" bson:"email,omitempty"`
	EmailVerified     string `json:"EmailVerified,omitempty" bson:"EmailVerified,omitempty"`
	ContactNo         string `json:"ContactNo,omitempty" bson:"ContactNo,omitempty"`
	PhoneVerified     string `json:"PhoneVerified,omitempty" bson:"PhoneVerified,omitempty"`
	Status            string `json:"Status,omitempty" bson:"Status,omitempty" `
	PasswordHash      string `json:"passwordHash,omitempty" bson:"passwordHash,omitempty"`
	TimestampCreated  int64  `json:"timestampCreated,omitempty" bson:"timestampCreated,omitempty"`
	TimestampModified int64  `json:"timestampModißfied,omitempty" bson:"timestampModified,omitempty"`
}

/*
func NewProfile(password string) *Profile {

	//func NewProfile(w http.ResponseWriter, r *http.Request) *Profile {
	u := Profile{}
	//fmt.Printf("This is NewProfile arg parameter print %v %v %v %v", Profilename, firstName, lastName, email, password)
	passwordHash := utility.SHA256OfString(u.PasswordHash)
	now := time.Now()
	unixTimestamp := now.Unix()
	b := Profile{u.CoEntity: u.CoEntity, Profilename: u.Profilename, FirstName: u.FirstName, LastName: u.LastName, Email: u.Email, ContactNo: u.ContactNo, PasswordHash: passwordHash, TimestampCreated: unixTimestamp}
	fmt.Printf("this is u %v", b)
	return &b
}
*/
