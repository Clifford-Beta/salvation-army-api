package staff

import (
	"salv_prj/model"
	"salv_prj/store"
)

type StaffService interface {
	AddStaff(staff model.Staff)(*model.Staff,error)
	RetrieveStaff(id int)(model.StaffResult, error)
	RetrieveAllStaff()([]*model.StaffResult,error)
	AddStaffRole(role model.StaffRole)(*model.StaffRole, error)
	RetrieveStaffRole(id int)(model.StaffRole,error)
	RetrieveAllRoles()([]*model.StaffRole,error)
	RecordBestPerformingStaff(teacher model.BestTeacher)(*model.BestTeacher,error)
	RetrieveBestPerformingStaff(from, to int)(model.BestStudentResult)
	RankStaffPerformance(from, to int)([]*model.BestStudentResult)
}

type Staffservice struct {}

func (Staffservice)AddStaff(staff model.Staff)(*model.Staff,error)  {
	staffStore := store.SqlStaffStore{store.Database}
	staff.Status = 1
	res := <-staffStore.CreateStaffMember(&staff)
	if res.Err != nil {
		return &model.Staff{},res.Err
	}
	return res.Data.(*model.Staff),nil
}

func (Staffservice)	RetrieveStaff(id int)(model.StaffResult, error)  {
	staffStore := store.SqlStaffStore{store.Database}
	res := <-staffStore.RetrieveStaffMember(id)
	if res.Err != nil {
		return model.StaffResult{},res.Err
	}
	return res.Data.(model.StaffResult),nil
}

func (Staffservice)	RetrieveAllStaff()([]*model.StaffResult,error)  {
	staffStore := store.SqlStaffStore{store.Database}
	res := <-staffStore.RetrieveAllStaffMembers()
	if res.Err != nil {
		return []*model.StaffResult{},res.Err
	}
	return res.Data.([]*model.StaffResult),nil
}

func (Staffservice) AddStaffRole(role model.StaffRole)(*model.StaffRole, error)  {
	staffStore := store.SqlStaffStore{store.Database}
	role.Status = 1
	res := <-staffStore.CreateStaffMemberRole(&role)
	if res.Err != nil {
		return &model.StaffRole{},res.Err
	}
	return res.Data.(*model.StaffRole),nil
}

func (Staffservice)	RetrieveStaffRole(id int)(model.StaffRole,error)  {
	staffStore := store.SqlStaffStore{store.Database}
	res := <-staffStore.RetrieveStaffMemberRole(id)
	if res.Err != nil {
		return model.StaffRole{},res.Err
	}
	return res.Data.(model.StaffRole),nil
}

func (Staffservice)	RetrieveAllRoles()([]*model.StaffRole,error)  {
	staffStore := store.SqlStaffStore{store.Database}
	res := <-staffStore.RetrieveAllStaffMemberRoles()
	if res.Err != nil {
		return []*model.StaffRole{},res.Err
	}
	return res.Data.([]*model.StaffRole),nil
}

//func (Staffservice)RecordBestPerformingStaff(teacher model.BestTeacher)(*model.BestTeacher,error)  {
//	staffStore := store.SqlStaffStore{store.Database}
//	teacher.Status = 1
//	res := <-staffStore.Create(&staff)
//	if res.Err != nil {
//		return &model.Staff{},res.Err
//	}
//	return res.Data.(*model.Staff),nil
//}