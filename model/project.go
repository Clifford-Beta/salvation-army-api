package model

import (
	"encoding/json"
	"io"
	"time"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Project struct {
	Id          int       `db:"project_id" json:"id"`
	School      int       `db:"school" json:"school"`
	Name        string    `db:"project_name" json:"name"`
	Description string    `db:"project_desc" json:"description"`
	Start       time.Time `db:"project_start" json:"start"`
	Duration    int       `db:"project_duration" json:"duration"`
	Progress    int       `db:"project_progress" json:"progress"`
	Status      int       `db:"project_status" json:"status"`
	TimeStamp   time.Time `db:"timestamp" json:"time_stamp"`
}

type ProjectResult struct {
	Id          int       `db:"id" json:"id"`
	School      string    `db:"school" json:"school"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Start       time.Time `db:"start" json:"start"`
	Duration    int       `db:"duration" json:"duration"`
	Progress    int       `db:"progress" json:"progress"`
	TimeStamp   time.Time `db:"time_stamp" json:"time_stamp"`
}


func (s Project) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required, validation.Length(3, 50), is.Alphanumeric),
		validation.Field(&s.Description, validation.Required, validation.Length(3, 200)),
		validation.Field(&s.School, validation.Required),
		validation.Field(&s.Start, validation.Required),
	)
}


func (o *Project) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ProjectFromJson(data io.Reader) *Project {
	decoder := json.NewDecoder(data)
	var o Project
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}
