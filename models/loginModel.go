package models

type LoginSt struct {
	Username string `json:"Username" bson:"Username"`
	Password string `json:"Password" bson:"Password"`
}

type SessionSt struct {
	Username string `json:"Username" bson:"Username"`
	Session  string `json:"Session" bson:"Session"`
}
