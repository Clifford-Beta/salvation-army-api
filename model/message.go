package model

import (
	"encoding/json"
	"io"
	"time"
	"github.com/go-ozzo/ozzo-validation"
)

type Message struct {
	Id         int       `db:"message_id" json:"id"`
	Title      string    `db:"message_title" json:"title"`
	Content    string    `db:"message_content" json:"content"`
	Attachment string    `db:"message_file" json:"attachment"`
	From       int    `db:"message_from" json:"from"`
	To         int    `db:"message_to" json:"to"`
	Status     int       `db:"message_status" json:"status"`
	TimeStamp  time.Time `db:"timestamp" json:"time_stamp"`
	DateSent   time.Time `db:"date_sent" json:"date_sent"`
}

type MessageResult struct {
	Id         int       `db:"message_id" json:"id"`
	Title      string    `db:"message_title" json:"title"`
	Content    string    `db:"message_content" json:"content"`
	Attachment string    `db:"message_file" json:"attachment"`
	From       string    `db:"message_from" json:"from"`
	To         string    `db:"message_to" json:"to"`
	Status     int       `db:"message_status" json:"status"`
	DateSent   time.Time `db:"date_sent" json:"date_sent"`
}

type FileType struct {
	Id          int       `db:"file_type_id" json:"id"`
	Name        string    `db:"file_type_name" json:"name"`
	Description string    `db:"file_type_desc" json:"description"`
	Store       string    `db:"file_type_store" json:"store"`
	Status      int       `db:"file_type_status" json:"status"`
	TimeStamp   time.Time `db:"timestamp" json:"time_stamp"`
}

type File struct {
	Id          int       `db:"file_id" json:"id"`
	Type        int       `db:"file_type" json:"type"`
	Name        string    `db:"file_name" json:"name"`
	Description string    `db:"file_description" json:"description"`
	Status      int       `db:"file_status" json:"status"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	TimeStamp   time.Time `db:"timestamp" json:"time_stamp"`
}

func (o Message) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.From, validation.Required),
		validation.Field(&o.To, validation.Required),
		validation.Field(&o.Title, validation.Required),
		validation.Field(&o.Content, validation.Required),
	)
}

func (o File) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Name, validation.Required, validation.Length(2, 75)),
		validation.Field(&o.Description, validation.Required, validation.Length(3, 200)),
		validation.Field(&o.Type, validation.Required, validation.Min(1)),
	)
}

func (o FileType) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Name, validation.Required, validation.Length(2, 75)),
		validation.Field(&o.Description, validation.Required, validation.Length(3, 200)),
		validation.Field(&o.Store, validation.Required, validation.Length(2, 30)),
	)
}

func (o *Message) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func MessageFromJson(data io.Reader) *Message {
	decoder := json.NewDecoder(data)
	var o Message
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *File) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func FileFromJson(data io.Reader) *File {
	decoder := json.NewDecoder(data)
	var o File
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *FileType) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func FileTypeFromJson(data io.Reader) *FileType {
	decoder := json.NewDecoder(data)
	var o FileType
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}
