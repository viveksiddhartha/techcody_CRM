package models

type CoEntity struct {
	UUID              string `json:"uuid" bson:"uuid"`
	CoEntityId        string `json:"coEntityId" bson:"coEntityId"`
	CompanyNm         string `json:"companyNm" bson:"companyNm"`
	AliasNm           string `json:"aliasNm" bson:"aliasNm"`
	State             string `json:"State" bson:"State"`
	Country           string `json:"Country" bson:"Country"`
	Email             string `json:"Email" bson:"Email"`
	SecretKey         string `json:"SecretKey" bson:"SecretKey"`
	TimestampCreated  int64  `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64  `json:"timestampModiÃŸfied" bson:"timestampModified"`
}
