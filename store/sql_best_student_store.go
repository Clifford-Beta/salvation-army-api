package store

import (
	"salvation-army-api/model"
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

func (s SqlBestStudentStore) Update(best *model.BestStudent) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if count, err := s.GetMaster().Update(best); err != nil {
			result.Err = model.NewLocAppError("SqlBestStudentStore.Update", "store.sql_best_student.update.updating.app_error", nil, "student_id="+strconv.Itoa(best.Id)+", "+err.Error())

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
		var best model.BestStudentResult
		err := s.master.SelectOne(&best, `select *
			 from best_student
			 where best_student_status = 1 and best_student_id=?`,id)

		if err != nil {
			result.Err = model.NewLocAppError("SqlBestStudentStore.Get", "store.sql_best_student.get.app_error", nil, err.Error())
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

func (s SqlBestStudentStore) GetBest(from, to int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var best model.BestStudentResult
		err := s.master.SelectOne(&best, `select best_student_id as id, school.school_name as school, best_student_name as name, best_student_class as
			 class, best_student_year as year, best_student_technique as technique, best_student_photo as photo,
			 max(best_student_mark) as mark,category.category_name as category, best_student.timestamp as time_stamp
			 from best_student
			 inner join school on best_student_school = school_id
			 inner join category on best_student.category = category_id
			 where best_student_status = 1`)
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

func (s SqlBestStudentStore) GetMany() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var bests []*model.BestStudentResult
		_, err := s.GetMaster().Select(&bests, `select best_student_id as id, school.school_name as school, best_student_name as name, best_student_class as
			 class, best_student_year as year, best_student_technique as technique, best_student_photo as photo,
			 best_student_mark as mark,category.category_name as category, best_student.timestamp as time_stamp
			 from best_student
			 inner join school on best_student_school = school_id
			 inner join category on best_student.category = category_id
			 where best_student_status = 1
			 order by mark desc`)
		if err != nil {
			result.Err = model.NewLocAppError("SqlBestStudentStore.GetMany", "store.sql_best_student .getmany.app_error", nil, err.Error())

		} else {
			if len(bests) == 0 {
				result.Err = model.NewLocAppError("SqlBestStudentStore.GetMany", "store.sql_best_student .getmany.app_error", nil, "No records found")

			}
			result.Data = bests
		}
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}
