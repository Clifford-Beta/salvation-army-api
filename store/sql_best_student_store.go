package store

import (
	"salv_prj/model"
	"strconv"
)

type SqlBestStudentStore struct {
	*SqlStore
}

func (s SqlBestStudentStore) Save(best *model.BestStudent) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(best); err != nil {

			result.Err = model.NewLocAppError("SqlBestStudentStore.Save", "store.sql_beststudent.save.app_error", nil, "best student="+best.Name+", "+err.Error())

		} else {
			//user.Sanitize()
			result.Data = best
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}


func (s SqlBestStudentStore) Delete(best *model.BestStudent) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Exec("Update best_student SET best_student_status=0 where best_student_id=?", best.Id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlBestStudentStore.Delete", "store.sql_best_student.delete.app_error", nil, "best_student_id="+strconv.Itoa(best.Id)+", "+err.Error())

		} else {
			result.Data = res
			//result.Err =
		}
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlBestStudentStore) Get(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var best model.BestStudent
		err := s.master.SelectOne(&best, "select * from best_student where best_student_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlBestStudentStore.Get", "store.sql_best_student.get.app_error", nil, "best student="+best.Name+", "+err.Error())
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


func (s SqlBestStudentStore)GetMany() StoreChannel  {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var bests [] *model.BestStudent
		_, err := s.GetMaster().Select(&bests, "SELECT * FROM best_student WHERE best_student_status=1")
		if err != nil {
			result.Err = model.NewLocAppError("SqlBestStudentStore.GetMany", "store.sql_best_student .getmany.app_error", nil, err.Error())

		}else {
			result.Data = bests
		}
		storeChannel<-result
		close(storeChannel)
	}()
	return storeChannel
}
