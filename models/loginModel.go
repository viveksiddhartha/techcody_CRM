package models

type LoginSt struct {
	Username string `json:"Username,omitempty" bson:"Username,omitempty"`
	Password string `json:"Password,omitempty" bson:"Password,omitempty"`
}

type SessionSt struct {
	Username string `json:"Username,omitempty" bson:"Username"`
	Session  string `json:"Session,omitempty" bson:"Session,omitempty"`
}
