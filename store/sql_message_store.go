package store

import (
	"salvation-army-api/model"
	"strconv"
)

type SqlMessageStore struct {
	*SqlStore
}

func (s SqlMessageStore) CreateMessage(message *model.Message) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(message); err != nil {

			result.Err = model.NewLocAppError("SqlMessageStore.CreateMessage", "store.sql_message.save.app_error", nil, err.Error())

		} else {
			//user.Sanitize()
			result.Data = message
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlMessageStore) Update(message *model.Message) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if count, err := s.GetMaster().Update(message); err != nil {
			result.Err = model.NewLocAppError("SqlMessageStore.Update", "store.sql_school.message.updating.app_error", nil, "user_id="+strconv.Itoa(message.Id)+", "+err.Error())

		}else{
			if count == 1 {
				result.Data = true
			}else{
				result.Data = false
			}

		}

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}


func (s SqlMessageStore) RetrieveMessage(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var message model.Message
		err := s.master.SelectOne(&message, "select * from message where message_id=?", id)
		if err != nil{
			result.Err = model.NewLocAppError("SqlMessageStore.Get", "store.sql_message.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = message
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlMessageStore) RetrieveAll(user int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var messages []model.MessageResult
		_, err := s.master.Select(&messages, `select message_id,message_title,message_content,message_file,username as message_from, message_status,date_sent
													from message
													inner join user on message.message_from = user_id
													where message_to = :user
													and message_status = :status`, map[string]interface{}{"status":1,"user":user})
		if err != nil {
			result.Err = model.NewLocAppError("SqlMessageStore.GetAll", "store.sql_message.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		} else {
			if len(messages) == 0 {
				result.Err = model.NewLocAppError("SqlMessageStore.RetrieveAll", "store.sql_message.retrieve_all.app_error", nil, "No records found")

			}
		}
		result.Data = messages
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}
