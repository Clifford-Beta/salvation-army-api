package store

import (
	"salv_prj/model"
	"time"
)

type SqlStaffStore struct {
	*SqlStore
}

type StaffResult struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Phone       string    `json:"phone" db:"phone"`
	Role        string    `json:"role" db:"role"`
	Photo       string    `json:"photo" db:"photo"`
	School      string    `json:"school" db:"school"`
	Title       string    `json:"title" db:"title"`
	Password    string    `json:"password" db:"password"`
	DateCreated time.Time `json:"date_created" db:"date_created"`
	TimeStamp   time.Time `json:"time_stamp" db:"timestamp"`
	Status      int       `json:"status" db:"status"`
}

func (s SqlStaffStore) CreateStaffMember(staff *model.Staff) StoreChannel  {
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
func (s SqlStaffStore) CreateStaffMemberRole(staff *model.StaffRole) StoreChannel  {
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
		var staff StaffResult
		err := s.master.SelectOne(&staff,
		"select staff.staff_id as id,staff.staff_name as name,staff.staff_phone as phone," +
			"staff.staff_email as email,staff.staff_photo as photo,staff.staff_title as title," +
			"staff.staff_status as status,staff.password as password, staff_role.staff_role_name as role" +
			",staff.date_created as date_created, staff.timestamp as timestamp,school.school_name as school " +
			"from `staff` " +
			"left join school on staff.school_id=school.school_id " +
			"left join staff_role on staff.staff_role=staff_role.staff_role_id " +
			"where staff.staff_id=?",id)

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
		var staff []StaffResult
		_,err := s.master.Select(&staff,
			"select staff.staff_id as id,staff.staff_name as name,staff.staff_phone as phone," +
				"staff.staff_email as email,staff.staff_photo as photo,staff.staff_title as title," +
				"staff.staff_status as status,staff.password as password, staff_role.staff_role_name as role" +
				",staff.date_created as date_created, staff.timestamp as timestamp,school.school_name as school " +
				"from `staff` " +
				"left join school on staff.school_id=school.school_id " +
				"left join staff_role on staff.staff_role=staff_role.staff_role_id " +
				"where staff.staff_status=?", 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlStaffStore.GetAll", "store.sql_staff.get.app_error", nil, err.Error())
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
func (s SqlStaffStore) RetrieveAllStaffMemberRoles() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var role []*model.StaffRole
		_,err := s.master.Select(&role, "select * from staff_role where staff_role_status=?", 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlStaffStore.GetRoles", "store.sql_staff_role.get.app_error", nil,err.Error())
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
