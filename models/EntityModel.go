package models

type CoEntity struct {
	UUID              string `json:"uuid,omitempty" bson:"uuid"`
	CoEntityId        string `json:"coEntityId,omitempty" bson:"coEntityId,omitempty"`
	CompanyNm         string `json:"companyNm,omitempty" bson:"companyNm,omitempty"`
	AliasNm           string `json:"aliasNm,omitempty" bson:"aliasNm,omitempty"`
	State             string `json:"State,omitempty" bson:"State,omitempty"`
	Country           string `json:"Country,omitempty" bson:"Country,omitempty"`
	Email             string `json:"Email,omitempty" bson:"Email,omitempty"`
	SecretKey         string `json:"SecretKey,omitempty" bson:"SecretKey,omitempty"`
	Status            string `json:"Status,omitempty" bson:"Status,omitempty"`
	Password          string `json:"Password,omitempty" bson:"Password,omitempty"`
	TimestampCreated  int64  `json:"timestampCreated,omitempty" bson:"timestampCreated,omitempty"`
	TimestampModified int64  `json:"timestampModified,omitempty" bson:"timestampModified,omitempty"`
}
