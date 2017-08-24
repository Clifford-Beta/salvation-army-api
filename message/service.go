package message

import (
	"salvation-army-api/model"
	"salvation-army-api/store"
	"time"
)

// StringService provides operations on strings.
type MessageService interface {
	Create(message model.Message) (*model.Message, error)
	GetOne(int) (model.Message, error)
	GetAll() (map[string][]model.Message, error)
	//Update(user model.User) (model.User, error)
}

type Messsageservice struct{}

func (Messsageservice) Create(message model.Message) (*model.Message, error) {
	msgStore := store.SqlMessageStore{store.Database}
	message.Status = 1
	message.DateSent = time.Now()
	me := <-msgStore.CreateMessage(&message)
	if me.Err != nil {
		return &model.Message{}, me.Err
	}
	return me.Data.(*model.Message), nil
}

func (Messsageservice) GetOne(id int) (model.Message, error) {
	msgStore := store.SqlMessageStore{store.Database}
	me := <-msgStore.RetrieveMessage(id)
	if me.Err != nil {
		return model.Message{}, me.Err
	}
	return me.Data.(model.Message), nil
}

func (Messsageservice) GetAll() (map[string][]model.Message, error) {
	msgStore := store.SqlMessageStore{store.Database}
	me := <-msgStore.RetrieveAll()
	if me.Err != nil {
		return map[string][]model.Message{"data": []model.Message{}}, me.Err
	}
	return map[string][]model.Message{"data": me.Data.([]model.Message)}, nil
}

//func (Userservice) Update(a,b int) (int, error) {
//	return a+b, nil
//}
