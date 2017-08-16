package store

import (
	"testing"
	"time"
	"salv_prj/model"
)

var staffStore = SqlStaffStore{Database}
func TestSqlStaffStore_CreateStaffMemberRole(t *testing.T) {
	staffrole := model.StaffRole{
		Id:2,
		Name:"Head Teacher",
		Description:"The primary administrator",
		TimeStamp:time.Now(),
		Status:1,
	}
	res := <- staffStore.CreateStaffMemberRole(&staffrole)
	if res.Err != nil {
		t.Errorf("CreateStaffMemberRole Failed with",res.Err)
	}
}

func TestSqlStaffStore_CreateStaffMember(t *testing.T) {
	staff := model.Staff{
		Name:"Zeph Adar",
		Email:"adarzeph@gmail.com",
		Phone:"0712379144",
		Role:2,
		Photo:"snisnyuegbe.jpg",
		School:1,
		Title:"H/T",
		Password:"1278nu8at723ow09742ue",
		DateCreated:time.Now(),
		TimeStamp:time.Now(),
		Status:1,
	}
	res := <- staffStore.CreateStaffMember(&staff)
	if res.Err != nil {
		t.Errorf("CreateStaffMember Failed with",res.Err)
	}
}

func TestSqlStaffStore_RetrieveStaffMember(t *testing.T) {
	res := <- staffStore.RetrieveStaffMember(7)
	if res.Err != nil{
		t.Errorf("RetrieveStaffMember failed with",res.Err)
	}
}



func TestSqlStaffStore_RetrieveAllStaffMemberRoles(t *testing.T) {
	res := <- staffStore.RetrieveAllStaffMemberRoles()
	if res.Err != nil{
		t.Errorf("RetrieveStaffMemberRoles failed with",res.Err)
	}
}

func TestSqlStaffStore_RetrieveAllStaffMembers(t *testing.T) {
	res := <- staffStore.RetrieveAllStaffMembers()
	if res.Err != nil{
		t.Errorf("RetrieveStaffMembers failed with",res.Err)
	}
}

