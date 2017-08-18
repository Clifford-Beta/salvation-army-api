package model

import (
	"time"
	"encoding/json"
	"io"
)

type Category struct {
	Id int `db:"category_id" json:"id"`
	Name string `db:"category_name" json:"name"`
	Description string `db:"category_description" json:"description"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
	Status	int	`db:"category_status" json:"status"`
}

type Tier struct {
	Id int `db:"tier_id" json:"id"`
	Name string `db:"tier_name" json:"name"`
	Description string `db:"tier_description" json:"description"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
	Status	int	`db:"tier_status" json:"status"`
}


func (o *Category) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func CategoryFromJson(data io.Reader) *Category {
	decoder := json.NewDecoder(data)
	var o Category
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}