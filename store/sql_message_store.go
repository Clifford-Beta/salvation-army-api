package store

import "salv_prj/model"

type SqlMessageStore struct {
	*SqlStore
}

func (s SqlMessageStore)CreateMessage(message *model.Message) StoreChannel {
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

func (s SqlMessageStore)RetrieveMessage(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var message model.Message
		err := s.master.SelectOne(&message,"select * from message where message_id=?",id)
		if err != nil {
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

func (s SqlMessageStore)RetrieveAll()StoreChannel  {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var messages []model.Message
		_,err := s.master.Select(&messages,"select * from message where message_status=?",1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlMessageStore.GetAll", "store.sql_message.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}else{
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
