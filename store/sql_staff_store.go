package store

import (
	"salvation-army-api/model"
	"strconv"
)

type SqlStaffStore struct {
	*SqlStore
}

func (s SqlStaffStore) CreateStaffMember(staff *model.Staff) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(staff); err != nil {

			result.Err = model.NewLocAppError("SqlStaffStore.CreateStaffMember", "store.sql_staff.save.app_error", nil, err.Error())

		} else {
			//user.Sanitize()
			result.Data = staff
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}
func (s SqlStaffStore) CreateStaffMemberRole(staff *model.StaffRole) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(staff); err != nil {

			result.Err = model.NewLocAppError("SqlStaffStore.CreateStaffMemberRole", "store.sql_staff.save.app_error", nil, err.Error())

		} else {
			//user.Sanitize()
			result.Data = staff
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlStaffStore) RetrieveStaffMember(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var staff model.Staff
		err := s.master.SelectOne(&staff,
			"select * from staff "+
				"where staff_status=1 and staff_id=?", id)
		//oldUserResult, err := s.GetMaster().Get(model.Staff{}, id)

		if err != nil {
			result.Err = model.NewLocAppError("SqlStaffStore.Get", "store.sql_category.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = staff

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}
func (s SqlStaffStore) Update(staff *model.Staff) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if sqlResult, err := s.GetMaster().Exec(
			`UPDATE
				staff
			SET
				staff_name = :Jina,
				staff_phone = :Phone,
				staff_email = :Email,
				staff_role = :Role,
				staff_photo = :Photo,
				staff_title = :Title,
				school_id = :School
			WHERE
				staff_id = :Id`,
			map[string]interface{}{
				"Id":             staff.Id,
				"Jina":             staff.Name,
				"Phone":      staff.Phone,
				"Email": 	staff.Email,
				"Role":         staff.Role,
				"Photo":         staff.Photo,
				"Title":         staff.Title,
				"School":         staff.School,
			}); err != nil {
			result.Err = model.NewLocAppError("SqlStaffStore.UpdateOptimistically",
				"store.sql_staff.update.app_error", nil, "id="+strconv.Itoa(staff.Id)+", "+err.Error())
		} else {
			rows, err := sqlResult.RowsAffected()

			if err != nil {
				result.Err = model.NewLocAppError("SqlStaffStore.UpdateStatus",
					"store.sql_staff.update.app_error", nil, "id="+strconv.Itoa(staff.Id)+", "+err.Error())
			} else {
				if rows == 1 {
					result.Data = true
				} else {
					result.Data = false
				}
			}
		}

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlStaffStore) Delete(staff *model.Staff) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Exec("Update staff SET staff_status=0 where staff_id=?", staff.Id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlStaffStore.Delete", "store.sql_staff.delete.app_error", nil, "staff_id="+strconv.Itoa(staff.Id)+", "+err.Error())

		} else {
			result.Data = res
		}
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}


func (s SqlStaffStore) RetrieveStaffMemberRole(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var role model.StaffRole
		err := s.master.SelectOne(&role, "select * from staff_role where staff_role_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlStaffStore.GetRole", "store.sql_staff_role.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		result.Data = role

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlStaffStore) RetrieveAllStaffMembers() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var staff []model.StaffResult
		_, err := s.master.Select(&staff,
			"select staff.staff_id as id,staff.staff_name as name,staff.staff_phone as phone,"+
				"staff.staff_email as email,staff.staff_photo as photo,staff.staff_title as title,"+
				"staff.staff_status as status, staff_role.staff_role_name as role"+
				",staff.date_created as date_created, staff.timestamp as timestamp,school.school_name as school "+
				"from `staff` "+
				"left join school on staff.school_id=school.school_id "+
				"left join staff_role on staff.staff_role=staff_role.staff_role_id "+
				"where staff.staff_status=?", 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlStaffStore.GetAll", "store.sql_staff.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		if len(staff) == 0 {
			result.Err = model.NewLocAppError("SqlStaffStore.RetrieveAll", "store.sql_staff.retrieve_all.app_error", nil, "No records found")

		}
		result.Data = staff

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}
func (s SqlStaffStore) RetrieveAllStaffMemberRoles() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var role []*model.StaffRole
		_, err := s.master.Select(&role, "select * from staff_role where staff_role_status=?", 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlStaffStore.GetRoles", "store.sql_staff_role.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		result.Data = role

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}
