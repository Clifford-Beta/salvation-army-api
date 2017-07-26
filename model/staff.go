package model

import (
	"time"
	"encoding/json"
	"io"
)

type StaffRole struct {
	Id int `db:"staff_role_id" json:"id"`
	Name string `db:"staff_role_name" json:"name"`
	Description string `db:"staff_role_description" json:"description"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
	Status	int	`db:"staff_role_status" json:"status"`
}


type Staff struct {
	Id int `db:"staff_id" json:"id"`
	Name string `db:"staff_name" json:"name"`
	Email string `db:"staff_email" json:"email"`
	Phone string `db:"staff_phone" json:"phone"`
	Role int `db:"staff_role" json:"role"`
	Photo string `db:"staff_photo" json:"photo"`
	School int `db:"school_id" json:"school"`
	Title string `db:"staff_title" json:"title"`
	Password string `db:"password" json:"password"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
	Status	int	`db:"staff_status" json:"status"`
}


func (o *Staff) ToJson() string {
	//set password to null
	o.Password = ""
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func StaffFromJson(data io.Reader) *Staff {
	decoder := json.NewDecoder(data)
	var o Staff
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *StaffRole) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func StaffRoleFromJson(data io.Reader) *StaffRole {
	decoder := json.NewDecoder(data)
	var o StaffRole
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}