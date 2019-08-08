package models

type LoginSt struct {
	Username string `json:"Username" bson:"Username"`
	Password string `json:"Password" bson:"Password"`
}
