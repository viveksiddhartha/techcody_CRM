package models

import "encoding/json"

type Contract struct {
	ContractID        string  `json:"ContractID,omitempty" bson:"ContractID,omitempty"`
	CoEntityID        string  `json:"CoEntityID,omitempty" bson:"CoEntityID,omitempty"`
	Version           float32 `json:"Version,omitempty" bson:"Version,omitempty"`
	EffectiveDate     string  `json:"EffectiveDate,omitempty" bson:"EffectiveDate,omitempty"`
	Status            string  `json:"Status,omitempty" bson:"Status,omitempty"`
	ContractType      string  `json:"ContractType,omitempty" bson:"ContractType,omitempty"`
	JsonObject        json.RawMessage
	Allocation        []Allocation `json:"Allocation,omitempty" bson:"Allocation,omitempty"`
	TimestampCreated  int64        `json:"timestampCreated,omitempty" bson:"timestampCreated,omitempty"`
	TimestampModified int64        `json:"timestampModified,omitempty" bson:"timestampModified,omitempty"`
}

type Allocation struct {
	ProfileName string  `json:"ProfileName,omitempty" bson:"ProfileName,omitempty"`
	Percentage  float32 `json:"Percentage,omitempty" bson:"Percentage,omitempty"`
	Relation    string  `json:"Relation,omitempty" bson:"Relation,omitempty"`
	Status      string  `json:"Status,omitempty" bson:"Status,omitempty"`
}

type AllocationList struct {
	UUID              string  `json:"uuid,omitempty" bson:"uuid"`
	ContractID        string  `json:"ContractID,omitempty" bson:"ContractID,omitempty"`
	CoEntityID        string  `json:"CoEntityID,omitempty" bson:"CoEntityID,omitempty"`
	Version           float32 `json:"Version,omitempty" bson:"Version,omitempty"`
	EffectiveDate     string  `json:"EffectiveDate,omitempty" bson:"EffectiveDate,omitempty"`
	ContractType      string  `json:"ContractType,omitempty" bson:"ContractType,omitempty"`
	ProfileName       string  `json:"ProfileName,omitempty" bson:"ProfileName,omitempty"`
	Percentage        float32 `json:"Percentage,omitempty" bson:"Percentage,omitempty"`
	Relation          string  `json:"Relation,omitempty" bson:"Relation,omitempty"`
	Status            string  `json:"Status,omitempty" bson:"Status,omitempty"`
	TimestampCreated  string  `json:"timestampCreated,omitempty" bson:"timestampCreated,omitempty"`
	TimestampModified string  `json:"timestampModified,omitempty" bson:"timestampModified,omitempty"`
}

/*

type Allocation struct {
	uuid          string    `json:"uuid"`
	ContractID    string    `json:"ContractID"`
	CoEntityID    string    `json:"CoEntityID"`
	Version       float64   `json:"Version"`
	EffectiveDate time.Time `json:"EffectiveDate"`
	Status        bool      `json:"Status"`
	ContractType  string    `json:"ContractType"`
	JsonObject    json.RawMessage
	Profile       []struct {
		ProfileName string  `json:"ProfileName"`
		Percentage  float64 `json:"Percentage"`
		Relation    string  `json:"Relation"`
		Status      string  `json:"Status,omitempty" bson:"Status,omitempty"`
	} `json:"Profile"`
	TimestampCreated  int64 `json:"timestampCreated,omitempty" bson:"timestampCreated,omitempty"`
	TimestampModified int64 `json:"timestampModified,omitempty" bson:"timestampModified,omitempty"`
}
*/
