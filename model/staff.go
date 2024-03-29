package model

import (
	"encoding/json"
	"io"
	"time"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type StaffRole struct {
	Id          int       `db:"staff_role_id" json:"id"`
	Name        string    `db:"staff_role_name" json:"name"`
	Description string    `db:"staff_role_description" json:"description"`
	TimeStamp   time.Time `db:"timestamp" json:"time_stamp"`
	Status      int       `db:"staff_role_status" json:"status"`
}

type Staff struct {
	Id          int       `db:"staff_id" json:"id"`
	Name        string    `db:"staff_name" json:"name"`
	Email       string    `db:"staff_email" json:"email"`
	Phone       string    `db:"staff_phone" json:"phone"`
	Role        int       `db:"staff_role" json:"role"`
	Photo       string    `db:"staff_photo" json:"photo"`
	School      int       `db:"school_id" json:"school"`
	Title       string    `db:"staff_title" json:"title"`
	Password    string    `db:"password" json:"password"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	TimeStamp   time.Time `db:"timestamp" json:"time_stamp"`
	Status      int       `db:"staff_status" json:"status"`
}

type StaffResult struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Phone       string    `json:"phone" db:"phone"`
	Role        string    `json:"role" db:"role"`
	Photo       string    `json:"photo" db:"photo"`
	School      string    `json:"school" db:"school"`
	Title       string    `json:"title" db:"title"`
	DateCreated time.Time `json:"date_created" db:"date_created"`
	TimeStamp   time.Time `json:"time_stamp" db:"timestamp"`
	Status      int       `json:"status" db:"status"`
}

func (s Staff) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required, validation.Length(3, 50), is.Alphanumeric),
		validation.Field(&s.Email, validation.Required, is.Email),
		validation.Field(&s.Phone, validation.Required, validation.Length(9, 15), is.Digit),
		validation.Field(&s.Role, validation.Required, validation.Min(1)),
		validation.Field(&s.School, validation.Required, validation.Min(1)),
	)
}

func (s StaffRole) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required, validation.Length(3, 50), is.Alphanumeric),
		validation.Field(&s.Description, validation.Required, validation.Length(5,150)),
	)
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


func (o *Staff) Sanitize() *Staff {
	//set password to null
	o.Password = ""
	return o
}
