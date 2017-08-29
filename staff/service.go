package staff

import (
	"salvation-army-api/model"
	"salvation-army-api/store"
	"time"
)

type StaffService interface {
	AddStaff(staff model.Staff) (*model.Staff, error)
	RetrieveStaff(id int) (model.StaffResult, error)
	RetrieveAllStaff() (map[string][]model.StaffResult, error)
	AddStaffRole(role model.StaffRole) (*model.StaffRole, error)
	RetrieveStaffRole(id int) (model.StaffRole, error)
	RetrieveAllRoles() (map[string][]*model.StaffRole, error)
	RecordBestPerformingStaff(teacher model.BestTeacher) (*model.BestTeacher, error)
	RecordBestPerformingStudent(student model.BestStudent) (*model.BestStudent, error)
	RetrieveBestPerformingStaff(from, to int) (model.BestTeacherResult, error)
	RetrieveBestPerformingStudent(from, to int) (model.BestStudentResult, error)
	RankStaffPerformance(from, to int) (map[string][]*model.BestTeacherResult, error)
	RankStudentPerformance(from, to int) (map[string][]*model.BestStudentResult, error)
}

type Staffservice struct{}

func (Staffservice) AddStaff(staff model.Staff) (*model.Staff, error) {
	staffStore := store.SqlStaffStore{store.Database}
	staff.Status = 1
	staff.TimeStamp = time.Now()
	if err:= staff.Validate(); err != nil {
		return &model.Staff{},err
	}
	res := <-staffStore.CreateStaffMember(&staff)
	if res.Err != nil {
		return &model.Staff{}, res.Err
	}
	return res.Data.(*model.Staff), nil
}

func (Staffservice) RetrieveStaff(id int) (model.StaffResult, error) {
	staffStore := store.SqlStaffStore{store.Database}
	res := <-staffStore.RetrieveStaffMember(id)
	if res.Err != nil {
		return model.StaffResult{}, res.Err
	}

	return res.Data.(model.StaffResult), nil
}

func (Staffservice) RetrieveAllStaff() (map[string][]model.StaffResult, error) {
	staffStore := store.SqlStaffStore{store.Database}
	res := <-staffStore.RetrieveAllStaffMembers()
	if res.Err != nil {
		return map[string][]model.StaffResult{"data": []model.StaffResult{}}, res.Err
	}
	//staffMembers :=
	return map[string][]model.StaffResult{"data": res.Data.([]model.StaffResult)}, nil
}

func (Staffservice) AddStaffRole(role model.StaffRole) (*model.StaffRole, error) {
	staffStore := store.SqlStaffStore{store.Database}
	role.Status = 1
	if err := role.Validate(); err != nil {
		return &model.StaffRole{},err
	}
	res := <-staffStore.CreateStaffMemberRole(&role)
	if res.Err != nil {
		return &model.StaffRole{}, res.Err
	}

	return res.Data.(*model.StaffRole), nil
}

func (Staffservice) RetrieveStaffRole(id int) (model.StaffRole, error) {
	staffStore := store.SqlStaffStore{store.Database}
	res := <-staffStore.RetrieveStaffMemberRole(id)
	if res.Err != nil {
		return model.StaffRole{}, res.Err
	}
	return res.Data.(model.StaffRole), nil
}

func (Staffservice) RetrieveAllRoles() (map[string][]*model.StaffRole, error) {
	staffStore := store.SqlStaffStore{store.Database}
	res := <-staffStore.RetrieveAllStaffMemberRoles()
	if res.Err != nil {
		return map[string][]*model.StaffRole{"data": []*model.StaffRole{}}, res.Err
	}
	return map[string][]*model.StaffRole{"data": res.Data.([]*model.StaffRole)}, nil
}

func (Staffservice) RecordBestPerformingStaff(teacher model.BestTeacher) (*model.BestTeacher, error) {
	staffStore := store.SqlBestTeacherStore{store.Database}
	teacher.Status = 1
	teacher.TimeStamp = time.Now()
	res := <-staffStore.Save(&teacher)
	if res.Err != nil {
		return &model.BestTeacher{}, res.Err
	}
	return res.Data.(*model.BestTeacher), nil
}

func (Staffservice) RecordBestPerformingStudent(teacher model.BestStudent) (*model.BestStudent, error) {
	staffStore := store.SqlBestStudentStore{store.Database}
	teacher.Status = 1
	teacher.TimeStamp = time.Now()
	res := <-staffStore.Save(&teacher)
	if res.Err != nil {
		return &model.BestStudent{}, res.Err
	}
	return res.Data.(*model.BestStudent), nil
}

func (Staffservice) RetrieveBestPerformingStaff(from, to int) (model.BestTeacherResult, error) {
	staffStore := store.SqlBestTeacherStore{store.Database}
	res := <-staffStore.Get(from, to)
	if res.Err != nil {
		return model.BestTeacherResult{}, res.Err
	}
	return res.Data.(model.BestTeacherResult), nil
}

func (Staffservice) RetrieveBestPerformingStudent(from, to int) (model.BestStudentResult, error) {
	staffStore := store.SqlBestStudentStore{store.Database}
	res := <-staffStore.Get(from, to)
	if res.Err != nil {
		return model.BestStudentResult{}, res.Err
	}
	return res.Data.(model.BestStudentResult), nil
}

func (Staffservice) RankStaffPerformance(from, to int) (map[string][]*model.BestTeacherResult, error) {
	staffStore := store.SqlBestTeacherStore{store.Database}
	res := <-staffStore.GetMany()
	if res.Err != nil {
		return map[string][]*model.BestTeacherResult{"data": []*model.BestTeacherResult{}}, res.Err
	}
	return map[string][]*model.BestTeacherResult{"data": res.Data.([]*model.BestTeacherResult)}, nil
}

func (Staffservice) RankStudentPerformance(from, to int) (map[string][]*model.BestStudentResult, error) {
	staffStore := store.SqlBestStudentStore{store.Database}
	res := <-staffStore.GetMany()
	if res.Err != nil {
		return map[string][]*model.BestStudentResult{"data": []*model.BestStudentResult{}}, res.Err
	}
	return map[string][]*model.BestStudentResult{"data": res.Data.([]*model.BestStudentResult)}, nil
}
