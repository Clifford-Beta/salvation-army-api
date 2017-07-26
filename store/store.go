package store

import (
	//"encoding/json"
	"errors"
	"github.com/go-gorp/gorp"
	"salv_prj/model"
	"encoding/json"
)

var Database = InitConnection()


type StoreResult struct {
	Data interface{}
	Err  *model.AppError
}

type StoreChannel chan StoreResult

type Store interface {
	User() UserStore

}


type UserStore interface {
	Save(user *model.User) StoreChannel
	Update(user *model.User) StoreChannel
	Delete(user *model.User) StoreChannel
	Get(id int64) StoreChannel
	GetByEmail(email string) StoreChannel
	GetByEmailAndPassword(email, password string) StoreChannel
	GetByPhoneAndPassword(phone, password string) StoreChannel
}


type userConverter struct{}

func (me userConverter) ToDb(val interface{}) (interface{}, error) {

	switch t := val.(type) {
	case model.StringMap:
		return model.MapToJson(t), nil
	case model.StringArray:
		return model.ArrayToJson(t), nil
		//case model.EncryptStringMap:
		//	return encrypt([]byte(utils.Cfg.SqlSettings.AtRestEncryptKey), model.MapToJson(t))
	case model.StringInterface:
		return model.StringInterfaceToJson(t), nil
	}

	return val, nil
}

func (me userConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	switch target.(type) {
	case *model.StringMap:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("store.sql.convert_string_map")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{Holder: new(string), Target: target, Binder: binder}, true
	case *model.StringArray:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("store.sql.convert_string_array")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{Holder: new(string), Target: target, Binder: binder}, true

	case *model.StringInterface:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("store.sql.convert_string_interface")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{Holder: new(string), Target: target, Binder: binder}, true
	}

	return gorp.CustomScanner{}, false
}
