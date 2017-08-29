package model

import (
	"encoding/json"
	"io"
	"time"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type School struct {
	Id             int       `db:"school_id" json:"id"`
	Name           string    `db:"school_name" json:"name"`
	Email          string    `db:"school_email" json:"email"`
	Phone          string    `db:"school_phone" json:"phone"`
	PostalAddress  string    `db:"school_postal_address" json:"postal_address"`
	Category       int       `db:"school_category" json:"category"`
	Logo           string    `db:"school_logo" json:"logo"`
	Location       string    `db:"school_location" json:"location"`
	Description    string    `db:"school_description" json:"description"`
	DateRegistered time.Time `db:"date_registered" json:"date_registered"`
	TimeStamp      time.Time `db:"timestamp" json:"time_stamp"`
	Status         int       `db:"school_status" json:"status"`
}

type SchoolResult struct {
	Id             int       `db:"id" json:"id"`
	Rank           int       `db:"rank" json:"rank"`
	Name           string    `db:"name" json:"name"`
	Email          string    `db:"email" json:"email"`
	Phone          string    `db:"phone" json:"phone"`
	PostalAddress  string    `db:"postal_address" json:"postal_address"`
	Category       string    `db:"category" json:"category"`
	Logo           string    `db:"logo" json:"logo"`
	Location       string    `db:"location" json:"location"`
	Description    string    `db:"description" json:"description"`
	DateRegistered time.Time `db:"date_registered" json:"date_registered"`
}

func (o *School) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (o School) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Name, validation.Required, validation.Length(3, 50), is.Alphanumeric),
		validation.Field(&o.Email, validation.Required, is.Email),
		validation.Field(&o.Phone, validation.Required, validation.Length(9, 15), is.Digit),
		validation.Field(&o.Location,  is.Alphanumeric),
		validation.Field(&o.PostalAddress, validation.Required,  is.Alphanumeric),
		validation.Field(&o.Category, validation.Required,  validation.Min(1)),
		validation.Field(&o.Description, validation.Required,  validation.Length(5,200)),
	)
}

func SchoolFromJson(data io.Reader) *School {
	decoder := json.NewDecoder(data)
	var o School
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}
