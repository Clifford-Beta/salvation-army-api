package model

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"time"
)

//	"time"
//	"encoding/json"
//	"io"
//	//"unicode"
//	"strings"
//	"fmt"
//	"regexp"
//	"net/http"
//	"golang.org/x/crypto/bcrypt"
//)

type User struct {
	Id       int       `db:"user_id" json:"id"`
	Name     string    `db:"username" json:"name"`
	Email    string    `db:"email" json:"email"`
	DateAdd  time.Time `db:"create_time" json:"date_add"`
	Password string    `db:"password" json:"password"`
	Status   int       `db:"status" json:"status"`
}

//func (u *User) IsValid() *AppError {
//	if !IsValidUsername(u.Name) {
//		return NewAppError("Insurer.IsValid", "model.user.is_valid.name.app_error", nil, "name="+u.Name, http.StatusBadRequest)
//	}
//	if !IsValidPhone(u.Phone) {
//		return NewAppError("Insurer.IsValid", "model.user.is_valid.phone.app_error", nil, "name="+u.Name, http.StatusBadRequest)
//	}
//	if !IsValidEmail(u.Email) {
//		return NewAppError("Insurer.IsValid", "model.user.is_valid.email.app_error", nil, "name="+u.Name, http.StatusBadRequest)
//	}
//	if !IsValidStatus(u.Status) {
//		return NewAppError("Insurer.IsValid", "model.user.is_valid.status.app_error", nil, "name="+u.Name, http.StatusBadRequest)
//	}
//
//	if !IsValidPassword(u.Password) {
//		return NewAppError("Insurer.IsValid", "model.insurer.is_valid.adress.app_error", nil, "name="+u.Name, http.StatusBadRequest)
//	}
//
//	return nil
//}
//
//func (u *User) Presave() {
//	u.Password = HashPassword(u.Password)
//}
//
//
//func IsValidInsurer(id int) bool{
//	if id > 0 {
//		return true
//	}
//	return false
//}
//
//func IsValidPassword(password string)  bool {
//	return true
//}
//
//// HashPassword generates a hash using the bcrypt.GenerateFromPassword
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

//
//// ComparePassword compares the hash
func ComparePassword(hash string, password string) bool {

	if len(password) == 0 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (o *User) ToJson() string {
	//set password to null
	o.Password = ""
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (o *User) Sanitize() *User {
	//set password to null
	o.Password = ""
	return o
}

func UserFromJson(data io.Reader) *User {
	decoder := json.NewDecoder(data)
	var o User
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

//
//var validUsernameChars = regexp.MustCompile(`^[a-z0-9\.\-_]+$`)
//
//var restrictedUsernames = []string{
//	"all",
//	"channel",
//	"matterbot",
//}
//
//func IsValidUsername(s string) bool {
//	if len(s) < USER_NAME_MIN_LENGTH || len(s) > USER_NAME_MAX_LENGTH {
//		return false
//	}
//
//	//if !validUsernameChars.MatchString(s) {
//	//	return false
//	//}
//	//
//	//if !unicode.IsLetter(rune(s[0])) {
//	//	return false
//	//}
//	//
//	//for _, restrictedUsername := range restrictedUsernames {
//	//	if s == restrictedUsername {
//	//		return false
//	//	}
//	//}
//
//	return true
//}
//
//func CleanUsername(s string) string {
//	s = strings.ToLower(strings.Replace(s, " ", "-", -1))
//
//	for _, value := range reservedName {
//		if s == value {
//			s = strings.Replace(s, value, "", -1)
//		}
//	}
//
//	s = strings.TrimSpace(s)
//
//	for _, c := range s {
//		char := fmt.Sprintf("%c", c)
//		if !validUsernameChars.MatchString(char) {
//			s = strings.Replace(s, char, "-", -1)
//		}
//	}
//
//	s = strings.Trim(s, "-")
//
//	if !IsValidUsername(s) {
//		s = "a" + NewId()
//	}
//
//	return s
//}
//
//func (u *User)Sanitize() {
//	u.Password = ""
//}
//
//
//func (o *User) ToJson() string {
//	//set password to null
//	o.Password = ""
//	b, err := json.Marshal(o)
//	if err != nil {
//		return ""
//	} else {
//		return string(b)
//	}
//}
//
//func UserFromJson(data io.Reader) *User {
//	decoder := json.NewDecoder(data)
//	var o User
//	err := decoder.Decode(&o)
//	if err == nil {
//		return &o
//	} else {
//		return nil
//	}
//}
//
//
//
//type UserReport struct {
//	Id int
//	Name string
//	Email string
//	Phone string
//	Role int
//	Insurer string
//	InsurerId int
//	DateAdd time.Time
//	Status	int
//}
