package model

import (
	"time"
	"encoding/json"
	"io"
)

type ExtraCurricular struct {
	Id int `db:"ext_curricular_id" json:"id"`
	Name string `db:"ext_curricular_name" json:"name"`
	Description string `db:"ext_curricular_desc" json:"description"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
	Status	int	`db:"ext_curricular_status" json:"status"`
}

type ExtraCurricularLevel struct {
	Id int `db:"ext_level_id" json:"id"`
	Name string `db:"ext_level_name" json:"name"`
	Description string `db:"ext_level_desc" json:"description"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
	Status	int	`db:"ext_level_status" json:"status"`
}

type ExtraCurricularActivity struct {
	Id int `db:"ext_activity_id" json:"id"`
	School int `db:"school" json:"school"`
	Level int `db:"level" json:"level"`
	Activity int `db:"activity" json:"activity"`
	Performance string `db:"ext_activity_performance" json:"performance"`
	Date time.Time `db:"date" json:"date"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
	Status	int	`db:"ext_activity_status" json:"status"`
}




func (o *ExtraCurricular) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ExtraCurricularFromJson(data io.Reader) *ExtraCurricular {
	decoder := json.NewDecoder(data)
	var o ExtraCurricular
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}


func (o *ExtraCurricularLevel) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ExtraCurricularLevelFromJson(data io.Reader) *ExtraCurricularLevel {
	decoder := json.NewDecoder(data)
	var o ExtraCurricularLevel
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}


func (o *ExtraCurricularActivity) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ExtraCurricularActivityFromJson(data io.Reader) *ExtraCurricularActivity {
	decoder := json.NewDecoder(data)
	var o ExtraCurricularActivity
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}
