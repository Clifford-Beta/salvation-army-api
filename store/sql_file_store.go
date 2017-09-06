package store

import (
	"salvation-army-api/model"
	"strconv"
)

type SqlFileStore struct {
	*SqlStore
}

func (s SqlFileStore) Create(inf *model.File) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(inf); err != nil {

			result.Err = model.NewLocAppError("SqlFileStore.Create", "store.sql_file.save.app_error", nil, err.Error())

		} else {
			//user.Sanitize()
			result.Data = inf
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlFileStore) Update(file *model.File) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if count, err := s.GetMaster().Update(file); err != nil {
			result.Err = model.NewLocAppError("SqlFileStore.Update", "store.sql_file.update.updating.app_error", nil, "file_id="+strconv.Itoa(file.Id)+", "+err.Error())

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


func (s SqlFileStore) CreateFileType(inf *model.FileType) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(inf); err != nil {

			result.Err = model.NewLocAppError("SqlFileStore.CreateType", "store.sql_file_type.save.app_error", nil, err.Error())

		} else {
			//user.Sanitize()
			result.Data = inf
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlFileStore) Delete(file *model.File) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Exec("Update file SET file_status=0 where file_id=?", file.Id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlFileStore.Delete", "store.sql_file.delete.app_error", nil, "file_id="+strconv.Itoa(file.Id)+", "+err.Error())

		} else {
			result.Data = res
		}
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlFileStore) RetrieveOne(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var inf model.File
		err := s.master.SelectOne(&inf, "select * from file where file_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlFileStore.Get", "store.sql_file.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = inf
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlFileStore) RetrieveOneType(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var inf model.FileType
		err := s.master.SelectOne(&inf, "select * from file_type where file_type_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlFileStore.GetType", "store.sql_file_type.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = inf
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlFileStore) RetrieveAll() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var infs []model.File
		_, err := s.master.Select(&infs, "select * from file where file_status=?", 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlFileStore.GetAll", "store.sql_file.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		} else {
			if len(infs) == 0 {
				result.Err = model.NewLocAppError("SqlFileStore.RetrieveAll", "store.sql_file .retrieve_all.app_error", nil, "No records found")

			}
		}
		result.Data = infs
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlFileStore) RetrieveAllTypes() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var infs []model.FileType
		_, err := s.master.Select(&infs, "select * from file_type where file_type_status=?", 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlFileStore.GetAllTypes", "store.sql_file_type.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = infs
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}
