package store

import (
	"salv_prj/model"
	"strconv"
)

type SqlBestTeacherStore struct {
	*SqlStore
}

func (s SqlBestTeacherStore) Save(best *model.BestTeacher) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(best); err != nil {

			result.Err = model.NewLocAppError("SqlBestTeacherStore.Save", "store.sql_best_teacher.save.app_error", nil, "best teacher="+best.Name+", "+err.Error())

		} else {
			//user.Sanitize()
			result.Data = best
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}


func (s SqlBestTeacherStore) Delete(best *model.BestTeacher) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Exec("Update best_teacher SET best_teacher_status=0 where best_teacher_id=?", best.Id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlBestteacherStore.Delete", "store.sql_best_teacher.delete.app_error", nil, "best_teacher_id="+strconv.Itoa(best.Id)+", "+err.Error())

		} else {
			result.Data = res
			//result.Err =
		}
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlBestTeacherStore) Get(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var best model.BestTeacher
		err := s.master.SelectOne(&best, "select * from best_teacher where best_teacher_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlBestTeacherStore.Get", "store.sql_best_teacher.get.app_error", nil, "best teacher="+best.Name+", "+err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		//user.Sanitize()
		result.Data = best

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}


func (s SqlBestTeacherStore)GetMany() StoreChannel  {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var bests [] *model.BestTeacher
		_, err := s.GetMaster().Select(&bests, "SELECT * FROM best_teacher WHERE best_teacher_status=1")
		if err != nil {
			result.Err = model.NewLocAppError("SqlBestTeacherStore.GetMany", "store.sql_best_teacher .getmany.app_error", nil, err.Error())

		}else {
			result.Data = bests
		}
		storeChannel<-result
		close(storeChannel)
	}()
	return storeChannel
}
