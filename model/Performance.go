package model

import (
	"encoding/json"
	"io"
	"time"
)

type SchoolPerformance struct {
	Id        int       `db:"s_performance_id" json:"id"`
	Mark      float64   `db:"s_performance_mark" json:"mark"`
	School    int       `db:"school" json:"school"`
	Category  int       `db:"s_performance_cat" json:"category"`
	Year      int       `db:"s_performance_year" json:"year"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
	Status    int       `db:"s_performance_status" json:"status"`
}

type SchoolPerformanceResult struct {
	Id             int       `db:"id" json:"id"`
	Mark           float64   `db:"mark" json:"mark"`
	School         string    `db:"school" json:"school"`
	Category       string    `db:"category" json:"category"`
	Location       string    `db:"location" json:"location"`
	Tier           string    `db:"tier" json:"tier"`
	Description    string    `db:"description" json:"description"`
	Year           int       `db:"year" json:"year"`
	DateRegistered time.Time `db:"date_registered" json:"date_registered"`
}

type BestStudent struct {
	Id        int       `db:"best_student_id" json:"id"`
	School    int       `db:"best_student_school" json:"school"`
	Name      string    `db:"best_student_name" json:"name"`
	Mark      float64   `db:"best_student_mark" json:"mark"`
	Class     string    `db:"best_student_class" json:"class"`
	Category  int       `db:"category" json:"category"`
	Age       int       `db:"best_student_age" json:"age"`
	Gender    string    `db:"best_student_gender" json:"gender"`
	Technique string    `db:"best_student_technique" json:"technique"`
	Photo     string    `db:"best_student_photo" json:"photo"`
	Year      int       `db:"best_student_year" json:"year"`
	Status    int       `db:"best_student_status" json:"status"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
}

type BestStudentResult struct {
	Id        int       `db:"id" json:"id"`
	School    string    `db:"school" json:"school"`
	Name      string    `db:"name" json:"name"`
	Mark      float64   `db:"mark" json:"mark"`
	Class     string    `db:"class" json:"class"`
	Category  string    `db:"category" json:"category"`
	Age       int       `db:"age" json:"age"`
	Gender    string    `db:"gender" json:"gender"`
	Technique string    `db:"technique" json:"technique"`
	Photo     string    `db:"photo" json:"photo"`
	Year      int       `db:"year" json:"year"`
	TimeStamp time.Time `db:"time_stamp" json:"time_stamp"`
}

type BestTeacher struct {
	Id        int       `db:"best_teacher_id" json:"id"`
	School    int       `db:"best_teacher_school" json:"school"`
	Name      string    `db:"best_teacher_name" json:"name"`
	Mark      float64   `db:"best_teacher_mark" json:"mark"`
	Class     string    `db:"best_teacher_class" json:"class"`
	Category  int       `db:"category" json:"category"`
	Gender    string    `db:"best_teacher_gender" json:"gender"`
	Technique string    `db:"best_teacher_technique" json:"technique"`
	Photo     string    `db:"best_teacher_photo" json:"photo"`
	Year      int       `db:"best_teacher_year" json:"year"`
	Status    int       `db:"best_teacher_status" json:"status"`
	TimeStamp time.Time `db:"timestamp" json:"time_stamp"`
}

type BestTeacherResult struct {
	Id        int       `db:"id" json:"id"`
	School    string    `db:"school" json:"school"`
	Name      string    `db:"name" json:"name"`
	Mark      float64   `db:"mark" json:"mark"`
	Class     string    `db:"class" json:"class"`
	Category  string    `db:"category" json:"category"`
	Gender    string    `db:"gender" json:"gender"`
	Technique string    `db:"technique" json:"technique"`
	Photo     string    `db:"photo" json:"photo"`
	Year      int       `db:"year" json:"year"`
	TimeStamp time.Time `db:"time_stamp" json:"time_stamp"`
}

func (o *SchoolPerformance) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func SchoolPerformaceFromJson(data io.Reader) *SchoolPerformance {
	decoder := json.NewDecoder(data)
	var o SchoolPerformance
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *BestStudent) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func BestStudentFromJson(data io.Reader) *BestStudent {
	decoder := json.NewDecoder(data)
	var o BestStudent
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *BestTeacher) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func BestTeacherFromJson(data io.Reader) *BestTeacher {
	decoder := json.NewDecoder(data)
	var o BestTeacher
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}
