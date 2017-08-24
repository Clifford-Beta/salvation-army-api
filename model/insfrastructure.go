package model

import (
	"encoding/json"
	"io"
	"time"
)

type Infrastructure struct {
	Id          int       `db:"infrastructure_id" json:"id"`
	School      int       `db:"school_id" json:"school"`
	Name        string    `db:"infrastructure_name" json:"name"`
	Type        int       `db:"infrastructure_type" json:"type"`
	Quantity    int       `db:"infrastructure_quantity" json:"quantity"`
	Description string    `db:"infrastructure_description" json:"description"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	TimeStamp   time.Time `db:"timestamp" json:"time_stamp"`
	Status      int       `db:"infrastructure_status" json:"status"`
}

type InfrastructureResult struct {
	Id          int       `db:"id" json:"id"`
	School      string       `db:"school" json:"school"`
	Name        string    `db:"name" json:"name"`
	Type        string       `db:"type" json:"type"`
	Quantity    int       `db:"quantity" json:"quantity"`
	Description string    `db:"description" json:"description"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

type InfrastructureType struct {
	Id          int       `db:"i_type_id" json:"id"`
	Name        string    `db:"i_type_name" json:"name"`
	Description string    `db:"i_type_description" json:"description"`
	TimeStamp   time.Time `db:"timestamp" json:"time_stamp"`
	Status      int       `db:"i_type_status" json:"status"`
}

func (o *Infrastructure) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func InfrastructureFromJson(data io.Reader) *Infrastructure {
	decoder := json.NewDecoder(data)
	var o Infrastructure
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *InfrastructureType) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func InfrastructureTypeFromJson(data io.Reader) *InfrastructureType {
	decoder := json.NewDecoder(data)
	var o InfrastructureType
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}
