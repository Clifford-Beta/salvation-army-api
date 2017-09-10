package store

import (
	"salvation-army-api/model"
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

func (s SqlBestTeacherStore) Update(best *model.BestTeacher) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if count, err := s.GetMaster().Update(best); err != nil {
			result.Err = model.NewLocAppError("SqlBestTeacherStore.Update", "store.sql_best_teacher.update.updating.app_error", nil, "teacher_id="+strconv.Itoa(best.Id)+", "+err.Error())

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
		//err := s.master.SelectOne(&best, `select best_teacher_id as id, school.school_name as school, best_teacher_name as name, best_teacher_class as
		//	 class, best_teacher_year as year, best_teacher_technique as technique, best_teacher_photo as photo,
		//	 max(best_teacher_mark) as mark,category.category_name as category, best_teacher.timestamp as time_stamp
		//	 from best_teacher
		//	 inner join school on best_teacher_school = school_id
		//	 inner join category on best_teacher.category = category_id
		//	 where best_teacher_status = 1`)
		//var best model.BestTeacher
		err := s.master.SelectOne(&best, `select * from best_teacher where best_teacher_status = 1 and best_teacher_id = ?`,id)

		if err != nil{
			result.Err = model.NewLocAppError("SqlBestTeacherStore.Get", "store.sql_best_teacher.get.app_error", nil, err.Error())
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

func (s SqlBestTeacherStore) GetBest(from,to int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var best model.BestTeacherResult
		err := s.master.SelectOne(&best, `select best_teacher_id as id, school.school_name as school, best_teacher_name as name, best_teacher_class as
			 class, best_teacher_year as year, best_teacher_technique as technique, best_teacher_photo as photo,
			 max(best_teacher_mark) as mark,category.category_name as category, best_teacher.timestamp as time_stamp
			 from best_teacher
			 inner join school on best_teacher_school = school_id
			 inner join category on best_teacher.category = category_id
			 where best_teacher_status = 1`)
		//var best model.BestTeacher
		//err := s.master.SelectOne(&best, `select * from best_teacher where best_teacher_status = 1 and best_teacher_id = ?`,id)

		if err != nil {
			result.Err = model.NewLocAppError("SqlBestTeacherStore.Get", "store.sql_best_teacher.get.app_error", nil, err.Error())
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

func (s SqlBestTeacherStore) GetMany() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var bests []*model.BestTeacherResult
		_, err := s.GetMaster().Select(&bests, `select best_teacher_id as id, school.school_name as school, best_teacher_name as name, best_teacher_class as
			 class, best_teacher_year as year, best_teacher_technique as technique, best_teacher_photo as photo,
			 best_teacher_mark as mark,category.category_name as category, best_teacher.timestamp as time_stamp
			 from best_teacher
			 inner join school on best_teacher_school = school_id
			 inner join category on best_teacher.category = category_id
			 where best_teacher_status = 1
			 order by mark desc`)
		if err != nil {
			result.Err = model.NewLocAppError("SqlBestTeacherStore.GetMany", "store.sql_best_teacher .getmany.app_error", nil, err.Error())

		} else {
			if len(bests) == 0 {
				result.Err = model.NewLocAppError("SqlBestTeacherStore.GetMany", "store.sql_best_teacher.getmany.app_error", nil, "No records found")

			}
		}
		result.Data = bests

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}
