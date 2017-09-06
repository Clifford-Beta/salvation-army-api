package message

import (
	"salvation-army-api/model"
	"salvation-army-api/store"
	"time"
)

type MessageService interface {
	Create(message model.Message) (*model.Message, error)
	GetOne(int) (model.Message, error)
	GetAll(user int) (map[string][]model.MessageResult, error)
	Update(msg model.Message) (bool, error)
	Delete(msg model.Message) (bool, error)
}

type Messsageservice struct{}

func (Messsageservice) Create(message model.Message) (*model.Message, error) {
	msgStore := store.SqlMessageStore{store.Database}
	message.Status = 1
	message.DateSent = time.Now()
	if err := message.Validate(); err != nil {
		return &model.Message{},err
	}
	me := <-msgStore.CreateMessage(&message)
	if me.Err != nil {
		return &model.Message{}, me.Err
	}
	return me.Data.(*model.Message), nil
}

func (Messsageservice) Update(msg model.Message) (bool, error) {
	projStore := store.SqlMessageStore{store.Database}
	me := <-projStore.Update(&msg)
	if me.Err != nil {
		return false, me.Err
	}
	return me.Data.(bool), nil
}

func (Messsageservice) Delete(msg model.Message) (bool, error) {
	projStore := store.SqlMessageStore{store.Database}
	me := <-projStore.Delete(&msg)
	if me.Err != nil {
		return false, me.Err
	}
	return true, nil
}


func (Messsageservice) GetOne(id int) (model.Message, error) {
	msgStore := store.SqlMessageStore{store.Database}
	me := <-msgStore.RetrieveMessage(id)
	if me.Err != nil {
		return model.Message{}, me.Err
	}
	return me.Data.(model.Message), nil
}

func (Messsageservice) GetAll(user int) (map[string][]model.MessageResult, error) {
	msgStore := store.SqlMessageStore{store.Database}
	me := <-msgStore.RetrieveAll(user)
	if me.Err != nil {
		return map[string][]model.MessageResult{"data": []model.MessageResult{}}, me.Err
	}
	return map[string][]model.MessageResult{"data": me.Data.([]model.MessageResult)}, nil
}

//func (Userservice) Update(a,b int) (int, error) {
//	return a+b, nil
//}
