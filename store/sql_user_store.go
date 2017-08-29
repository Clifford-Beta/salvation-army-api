package store

import (
	"log"
	"net/http"
	"salvation-army-api/model"
	"strconv"
	"fmt"
)

type SqlUserStore struct {
	*SqlStore
}

func (s SqlUserStore) Save(user *model.User) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//user.Presave()
		//if result.Err = user.IsValid(); result.Err != nil {
		//	storeChannel <- result
		//	close(storeChannel)
		//	return
		//}

		if err := s.GetMaster().Insert(user); err != nil {
			if IsUniqueConstraintError(err.Error(), []string{"Email", "users_email_key", "idx_users_email_unique"}) {
				result.Err = model.NewAppError("SqlUserStore.Save", "store.sql__user.save.email_exists.app_error", nil, err.Error(), http.StatusBadRequest)
			} else if IsUniqueConstraintError(err.Error(), []string{"Phone", "users_phone_key", "idx_users_phone_unique"}) {
				result.Err = model.NewAppError("SqlUserStore.Save", "store.sql__user.save.phone_exists.app_error", nil, err.Error(), http.StatusBadRequest)
			} else {
				result.Err = model.NewLocAppError("SqlInsurerUserStore.Save", "store.sql_insurer_user.save.app_error", nil, "user="+user.Name+", "+err.Error())
			}
		} else {
			//user.Sanitize()
			result.Data = user
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlUserStore) Update(user *model.User) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//oldUserResult, err := s.GetMaster().Get(model.User{}, user.Id)
		if oldUserResult, err := s.GetMaster().Get(model.User{}, user.Id); err != nil {
			result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.finding.app_error", nil, "user_id="+strconv.Itoa(user.Id))

		} else {
			oldUser := oldUserResult.(*model.User)
			user.DateAdd = oldUser.DateAdd
			user.Password = model.HashPassword(user.Password)
			if count, err := s.GetMaster().Update(user); err != nil {
				if IsUniqueConstraintError(err.Error(), []string{"Email", "users_email_key", "idx_user_email_unique"}) {
					result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.email_taken.app_error", nil, "user_id="+strconv.Itoa(user.Id)+", "+err.Error())
				} else if IsUniqueConstraintError(err.Error(), []string{"Phone", "users_phone_key", "idx_users_phone_unique"}) {
					result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.phone_taken.app_error", nil, "user_id="+strconv.Itoa(user.Id)+", "+err.Error())
				} else {
					result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.updating.app_error", nil, "user_id="+strconv.Itoa(user.Id)+", "+err.Error())
				}
			} else if count != 1 {
				result.Err = model.NewLocAppError("SqlInsurerUserStore.Update", "store.sql_insurer_user.update.app_error", nil, fmt.Sprintf("user_id=%v, count=%v", user.Id, count))
			} else {
				result.Data = [2]*model.User{user, oldUser}
			}
		}

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlUserStore) Delete(user *model.User) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Exec("Update insurer_user SET status=0 where insurer_user_id=?", user.Id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlInsurerUserStore.Delete", "store.sql_insurer_user.delete.app_error", nil, "user_id="+strconv.Itoa(user.Id)+", "+err.Error())

		} else {
			result.Data = res
			//result.Err =
		}
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlUserStore) Get(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var user model.User
		err := s.master.SelectOne(&user, "select * from user where user_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlUserStore.Get", "store.sql_user.get.app_error", nil, "user="+user.Name+", "+err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		//user.Sanitize()
		result.Data = user

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

//func (s SqlUserStore) GetByEmail(email string) StoreChannel {
//	storeChannel := make(StoreChannel, 1)
//	go func() {
//		result := StoreResult{}
//		var user model.User
//		err := s.master.Select(&user, "select * from insurer_user where email=?", email)
//		fmt.Println("Error", err)
//		if err != nil {
//			result.Err = model.NewLocAppError("SqlUserStore.GetByEmail", "store.sql_user.get.app_error", nil, "email="+email+", "+err.Error())
//			storeChannel <- result
//			close(storeChannel)
//			return
//		}
//
//		user.Sanitize()
//		result.Data = user
//
//		storeChannel <- result
//		close(storeChannel)
//	}()
//	return storeChannel
//}
//func (s SqlUserStore) GetByPhone(phone string) StoreChannel {
//	storeChannel := make(StoreChannel, 1)
//	go func() {
//		result := StoreResult{}
//		//pl := model.NewUserList()
//		var user model.User
//		err := s.master.SelectOne(&user, "select * from insurer_user where phone=?", phone)
//		if err != nil {
//			result.Err = model.NewLocAppError("SqlUserStore.GetByPhone", "store.sql_user.get.app_error", nil, "user ="+phone+", "+err.Error())
//			storeChannel <- result
//			close(storeChannel)
//			return
//		}
//
//		user.Sanitize()
//		result.Data = user
//
//		storeChannel <- result
//		close(storeChannel)
//	}()
//	return storeChannel
//}
//
//func (s SqlUserStore) GetByPhoneAndPassword(phone, password string) StoreChannel {
//	storeChannel := make(StoreChannel, 1)
//	go func() {
//		result := StoreResult{}
//		//pl := model.NewUserList()
//		var user model.User
//
//		err := s.master.SelectOne(&user, "SELECT * from insurer_user WHERE phone= :phone", map[string]interface{}{"phone": phone})
//		if err != nil {
//			result.Err = model.NewLocAppError("SqlUserStore.GetByPhoneAndPassword", "store.sql_user.get.app_error", nil, "user ="+phone+", "+err.Error())
//			storeChannel <- result
//			close(storeChannel)
//			return
//		}
//		if !model.ComparePassword(user.Password, password) {
//			result.Err = model.NewLocAppError("SqlUserStore.GetByPhoneAndPassword", "store.sql_user.get.app_error", nil,
//				"user ="+phone+":, The password provided does not match your current password")
//			storeChannel <- result
//			close(storeChannel)
//			return
//		}
//		user.Sanitize()
//		result.Data = user
//
//		storeChannel <- result
//		close(storeChannel)
//	}()
//	return storeChannel
//}

func (s SqlUserStore) GetByEmailAndPassword(email, password string) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var user model.User
		err := s.master.SelectOne(&user, "SELECT * from user WHERE email= :email", map[string]interface{}{"email": email})
		//err := s.master.SelectOne(&user, "select * from user where email = :email and password = 12345",map[string]interface{}{"email": email,"password": password})
		log.Println("This is the password provided", password, email)
		if err != nil {
			result.Err = model.NewLocAppError("SqlUserStore.GetByEmailAndPassword", "store.sql_user.get.app_error", nil, "user ="+email+", "+err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		log.Println("This is the user password", user.Password, "And this is the provided password", model.HashPassword(password))
		if !model.ComparePassword(user.Password, password) {
			result.Err = model.NewLocAppError("SqlUserStore.GetEmailAndPassword", "store.sql_user.get.app_error", nil,
				"user ="+email+":, The password provided does not match your current password")
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		user.Sanitize()
		result.Data = user

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlUserStore) GetMany() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var users []*model.User
		_, err := s.GetMaster().Select(&users, "SELECT * FROM user WHERE status=1")
		if err != nil {
			result.Err = model.NewLocAppError("SqlUsertore.GetMany", "store.sql_insurer_user.getmany.app_error", nil, err.Error())

		} else {
			for _, user := range users {
				user.Sanitize()
			}
			if len(users) == 0 {
				result.Err = model.NewLocAppError("SqlUserStore.GetMany", "store.sql_user.get_many.app_error", nil, "No records found")

			}
			result.Data = users
		}
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

//func (s SqlUserStore)GetManyByInsurer(id int) StoreChannel  {
//	storeChannel := make(StoreChannel, 1)
//	go func() {
//		result := StoreResult{}
//		var userreport [] *model.UserReport
//		_, err := s.GetMaster().Select(&userreport, `SELECT insurer_user.insurer_user_id as id,
//					 insurer_user.name as name, insurer_user.phone as phone,
//					insurer_user.email as email, insurer_user.date_time as dateadd,
//					insurer_user.status as status, insurer_user.role as role,
//					insurer.name as insurer, insurer.insurer_id as insurerid FROM insurer_user
//					INNER JOIN insurer
//					ON insurer_user.insurer_id=insurer.insurer_id
//					WHERE insurer_user.insurer_id = :id`,map[string]interface{}{"id": id})
//		if err != nil {
//			result.Err = model.NewLocAppError("SqlUsertore.GetMany", "store.sql_insurer_user.getmany.app_error", nil, err.Error())
//
//		}else {
//
//			result.Data = userreport
//		}
//		storeChannel<-result
//		close(storeChannel)
//	}()
//	return storeChannel
//}

//`SELECT insurer_user.insurer_user_id as id,
// insurer_user.name as name, insurer_user.phone as phone,
//insurer_user.email as email, insurer_user.date_time as dateregistered,
//insurer_user.status as status, insurer_user_role.description as role,
//insurer.name as insurer FROM insurer_user
//INNER JOIN insurer_user_role
//OUTER JOIN insurer
//ON insurer_user.role=insurer_user_role.insurer_user_role_id and
//insurer_user.insurer_id=insurer.insurer_id
//WHERE insurer_id = :id`,map[string]interface{}{"id": id})
