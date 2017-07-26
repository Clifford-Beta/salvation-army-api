package model

//import (
//	"encoding/json"
//	"io"
//	"net/http"
//	"go_private_api/store/utils"
//)
//
//const (
//	PHONE_NO_MAX_LENGTH = 14
//	PHONE_NO_MIN_LENGTH = 9
//	MESSAGE_MIN_LENGTH  = 0
//	MESSAGE_MAX_LENGTH  = 161
//)
//
//type SendSmsStruct struct {
//	Phone   string
//	Message string
//}
//
//func (r SendSmsStruct) IsValid() *utils.APIError {
//	if !IsValidPhone(r.Phone) {
//		return utils.NewAPIError("SendSmsStruct.IsValid", "sendyit.private_api.model.is_valid.phone", "phone="+r.Phone, http.StatusBadRequest)
//	}
//	if !IsValidMessage(r.Message) {
//		return utils.NewAPIError("SendSmsStruct.IsValid", "sendyit.private_api.model.is_valid.message", "phone="+r.Phone, http.StatusBadRequest)
//	}
//	return nil
//}
//
//func IsValidPhone(p string) bool {
//	if len(p) > PHONE_NO_MIN_LENGTH && len(p) < PHONE_NO_MAX_LENGTH {
//		return false
//	}
//	return true
//}
//
//func IsValidMessage(m string) bool {
//	if len(m) > MESSAGE_MIN_LENGTH && len(m) < MESSAGE_MAX_LENGTH {
//		return false
//	}
//	return true
//}
//
//func (r *SendSmsStruct) ToJSON() string {
//	w, err := json.Marshal(r)
//	if err != nil {
//		return ""
//	} else {
//		return string(w)
//	}
//}
//
//func SendSmsFromJson(data io.Reader) (*SendSmsStruct,error) {
//	decoder := json.NewDecoder(data)
//	var o SendSmsStruct
//	err := decoder.Decode(&o)
//	return &o,err
//}
