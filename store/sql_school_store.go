package store

import (
	"salv_prj/model"
	"net/http"
	"strconv"
)

type SqlSchoolStore struct {
	*SqlStore
}

func (s SqlSchoolStore) Save(school *model.School) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//user.Presave()
		//if result.Err = user.IsValid(); result.Err != nil {
		//	storeChannel <- result
		//	close(storeChannel)
		//	return
		//}

		if err := s.GetMaster().Insert(school); err != nil {
			if IsUniqueConstraintError(err.Error(), []string{"Email", "school_email_key", "idx_users_email_unique"}) {
				result.Err = model.NewAppError("SqlSchoolStore.Save", "store.sql_school.save.email_exists.app_error", nil, err.Error(), http.StatusBadRequest)
			} else if IsUniqueConstraintError(err.Error(), []string{"Phone", "school__phone_key", "idx_school_phone_unique"}) {
				result.Err = model.NewAppError("SqlSchoolStore.Save", "store.sql__school.save.phone_exists.app_error", nil, err.Error(), http.StatusBadRequest)
			} else {
				result.Err = model.NewLocAppError("SqlSchoolStore.Save", "store.sql_school.save.app_error", nil, "school="+school.Name+", "+err.Error())
			}
		} else {
			//user.Sanitize()
			result.Data = school
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

//func (s SqlUserStore) Update(user *model.User) StoreChannel {
//	storeChannel := make(StoreChannel, 1)
//	go func() {
//		result := StoreResult{}
//		if result.Err = user.IsValid(); result.Err != nil {
//			storeChannel <- result
//			close(storeChannel)
//			return
//		}
//		oldUserResult, err := s.GetMaster().Get(model.User{}, user.Id)
//		if err != nil || oldUserResult == nil {
//			result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.finding.app_error", nil, "user_id="+strconv.Itoa(user.Id))
//			//} else if oldInsurerResult == nil {
//			//	result.Err = model.NewLocAppError("SqlInsurerStore.Update", "store.sql_insurer.update.find.app_error", nil, "insurer_id="+strconv.Itoa(insurer.Id))
//		} else {
//			oldUser := oldUserResult.(*model.User)
//			user.DateAdd = oldUser.DateAdd
//			user.Password = model.HashPassword(user.Password)
//			if count, err := s.GetMaster().Update(user); err != nil {
//				if IsUniqueConstraintError(err.Error(), []string{"Email", "users_email_key", "idx_user_email_unique"}) {
//					result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.email_taken.app_error", nil, "user_id="+strconv.Itoa(user.Id)+", "+err.Error())
//				} else if IsUniqueConstraintError(err.Error(), []string{"Phone", "users_phone_key", "idx_users_phone_unique"}) {
//					result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.phone_taken.app_error", nil, "user_id="+strconv.Itoa(user.Id)+", "+err.Error())
//				} else {
//					result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.updating.app_error", nil, "user_id="+strconv.Itoa(user.Id)+", "+err.Error())
//				}
//			} else if count != 1 {
//				result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.app_error", nil, fmt.Sprintf("user_id=%v, count=%v", user.Id, count))
//			} else {
//				result.Data = [2]*model.User{user, oldUser}
//			}
//		}
//
//		storeChannel <- result
//		close(storeChannel)
//	}()
//	return storeChannel
//}

func (s SqlSchoolStore) Delete(school *model.School) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Exec("Update school SET status=0 where school_id=?", school.Id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlSchoolStore.Delete", "store.sql_school.delete.app_error", nil, "user_id="+strconv.Itoa(school.Id)+", "+err.Error())

		} else {
			result.Data = res
			//result.Err =
		}
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlSchoolStore) Get(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var school model.School
		err := s.master.SelectOne(&school, "select * from school where school_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlSchoolStore.Get", "store.sql_school.get.app_error", nil, "school="+school.Name+", "+err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		//user.Sanitize()
		result.Data = school

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}


func (s SqlSchoolStore)GetMany() StoreChannel  {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var schools [] *model.School
		_, err := s.GetMaster().Select(&schools, "SELECT * FROM school WHERE school_status=1")
		if err != nil {
			result.Err = model.NewLocAppError("SqlUsertore.GetMany", "store.sql_school.getmany.app_error", nil, err.Error())

		}else {
			result.Data = schools
		}
		storeChannel<-result
		close(storeChannel)
	}()
	return storeChannel
}