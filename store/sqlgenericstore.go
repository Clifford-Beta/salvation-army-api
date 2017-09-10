package store

import (
	"fmt"
	"reflect"
	"salvation-army-api/model"
	"time"
)

type SqlGenericStore struct {
	*SqlStore
}

type Returned interface {
	Test()
}
type Insurer struct {
	Id             int       `db:"insurer_id" `
	Name           string    `db:"name" `
	Email          string    `db:"email"`
	Phone          string    `db:"phone" `
	Status         int       `db:"status" `
	Company_Status int       `db:"company_status" `
	Address        string    `db:"address" `
	DateSignedUp   time.Time `db:"date_time" `
}

func (m Insurer) Test() {}

func (s SqlGenericStore) Select(params map[string]interface{}) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.New()
		var insurer Insurer
		err := s.master.SelectOne(&insurer, "select * from insurer where insurer_id=:id", params)
		fmt.Println("Error", err)
		if err != nil {
			result.Err = model.NewLocAppError("SqlInsurerStore.Get", "store.sql_insurer.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddInsurer(&insurer)

		result.Data = insurer

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlGenericStore) Run(query string) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	var insurer Insurer
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Select(&insurer, query)
		var params map[string]interface{}

		if err != nil {
			params["query"] = query
			result.Err = model.NewLocAppError("SqlGenericStore.Run", "store.sql_genericr.run.app_error", params, err.Error())
		} else {
			result.Data = res
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlGenericStore) Get(params map[string]interface{}, obj Returned) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.New()
		var insurer = obj
		fmt.Println(insurer, reflect.TypeOf(obj))
		err := s.master.SelectOne(&insurer, "select * from insurer where insurer_id=:id", params)
		fmt.Println("Error", err)
		if err != nil {
			result.Err = model.NewLocAppError("SqlInsurerStore.Get", "store.sql_insurer.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddInsurer(&insurer)

		result.Data = insurer

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}
