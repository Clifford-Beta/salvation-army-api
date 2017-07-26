package model

import (
	"testing"
	"time"

	"fmt"
)

func TestUser_ToJson(t *testing.T) {
	user := User{Id:1,
		Name:"Clifford Beta",
		Email:"betaclifford@gmail.com",
		Password:"126812ndst617281uj2e13",
		Status:1,
		DateAdd:time.Now(),
	}
	userjson := user.ToJson();
	if userjson == ""{
		t.Errorf("Failed on comnverting user to JSON")
	}
	fmt.Println("This is the resulting JSON",userjson)
}


func TestStaff_ToJson(t *testing.T) {
	staff := Staff{Id:1,
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
	staffjson := staff.ToJson();
	if staffjson == ""{
		t.Errorf("Failed on comnverting user to JSON")
	}
	fmt.Println("This is the resulting JSON",staffjson)
}

func TestStaffRole_ToJson(t *testing.T) {
	staffrole := StaffRole{
		Id:2,
		Name:"Head Teacher",
		Description:"The primary administrator",
		TimeStamp:time.Now(),
		Status:1,
	}
	staffrolejson := staffrole.ToJson();
	if staffrolejson == ""{
		t.Errorf("Failed on comnverting user to JSON")
	}
	fmt.Println("This is the resulting JSON",staffrolejson)
}

func TestSchool_ToJson(t *testing.T) {
	school := School{
		Id:1,
		Name:"Thika School",
		Email:"thika@slvarmy.com",
		Phone:"0912687904",
		PostalAddress:"1812912-00100",
		Category:2,
		Logo:"absiujmwiuhndowd.jpg",
		Location:"Thika Town",
		Description:"A school for the blind",
		DateRegistered:time.Now(),
		Status:1,
	}
	schooljson := school.ToJson()
	if schooljson == ""{
		t.Errorf("Failed on comnverting user to JSON")

	}
	fmt.Println("This is the resulting JSON",schooljson)

}

func TestProject_ToJson(t *testing.T) {
	proj := Project{
		Id:1,
		School:1,
		Name:"Road",
		Description:"Roads to all hostels in the school",
		Start:time.Now(),
		Duration:100,
		Progress:0,
		Status:1,
		TimeStamp:time.Now(),
	}
	projson := proj.ToJson()
	if projson == "" {
		t.Errorf("Failed on comnverting project to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)


}

func TestSchoolPerformace_ToJson(t *testing.T) {
	scp := SchoolPerformance{
		Id:1,
		Mark:87,
		School:1,
		Category:2,
		Year:2016,
		TimeStamp:time.Now(),
		Status: 1,
	}
	projson := scp.ToJson()
	if projson == "" {
		t.Errorf("Failed on comnverting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)

}

func TestBestStudent_ToJson(t *testing.T) {
	stud := BestStudent{
		Id:1,
		School:1,
		Name:"Clifford Beta",
		Class:"Form 4A",
		Category:1,
		Age:21,
		Gender:"M",
		Mark:98,
		Technique:"Nothing much",
		Photo:"tiosjmidusyhwmeiwoe.jpg",
		Status:1,
	}
	projson := stud.ToJson()
	if projson == "" {
		t.Errorf("Failed on comnverting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestBestTeacher_ToJson(t *testing.T) {
	stud := BestTeacher{
		Id:1,
		School:1,
		Name:"Clifford Beta",
		Class:"Form 4A",
		Category:1,
		Gender:"M",
		Mark:89,
		Technique:"Nothing much",
		Photo:"tiosjmidusyhwmeiwoe.jpg",
		Status:1,
	}
	projson := stud.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestMessage_ToJson(t *testing.T) {
	msg := Message{
		Id:1,
		Title:"Test",
		Content:"This is a test message",
		Attachment:"",
		From:"admin@system.com",
		To:"user@system.com",
		Status:1,
		TimeStamp:time.Now(),
		DateSent:time.Now(),
	}
	projson := msg.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestFileType_ToJson(t *testing.T) {
	fltp := FileType{
		Id:1,
		Name:"Memo",
		Description:"These are memos",
		Store:"Memoir",
		Status:1,
		TimeStamp:time.Now(),

	}
	projson := fltp.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestFile_ToJson(t *testing.T) {
	fl := File{
		Id:1,
		Type:1,
		Name:"memo_1_2016.pdf",
		Description:"First Memo of 2016",
		DateCreated:time.Now(),
		TimeStamp:time.Now(),

	}
	projson := fl.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestInfrastructure_ToJson(t *testing.T) {
	inf := Infrastructure{
		Id:1,
		School:1,
		Name:"Computer",
		Type:1,
		Quantity:100,
		Description:"Desktop computers",
		DateCreated:time.Now(),
		TimeStamp:time.Now(),
		Status:1,

	}
	projson := inf.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestInfrastructureType_ToJson(t *testing.T) {
	inft := InfrastructureType{
		Id:1,
		Name:"Electronics",
		Description:"Electricl Appliances",
		TimeStamp:time.Now(),
		Status:1,
	}
	projson := inft.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestExtraCurricularLevel_ToJson(t *testing.T) {
	extlvl := ExtraCurricularLevel{
		Id:1,
		Name:"Nationals",
		Description:"Nationwide competitions",
		TimeStamp:time.Now(),
		Status:1,

	}
	projson := extlvl.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestExtraCurricular_ToJson(t *testing.T) {
	extcr := ExtraCurricular{
		Id:1,
		Name:"Drama",
		Description:"Acting and drama festivals",
		TimeStamp:time.Now(),
		Status:1,
	}
	projson := extcr.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestExtraCurricularActivity_ToJson(t *testing.T) {
	extact := ExtraCurricularActivity{
		Id:1,
		School:1,
		Level:1,
		Activity:1,
		Performance:"First Runners Up",
		Date:time.Now(),
		TimeStamp:time.Now(),
		Status:1,

	}
	projson := extact.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}

func TestCategory_ToJson(t *testing.T) {
	cat := Category{
		Id:1,
		Name:"Overall",
		Description:"Cumulative Performance",
		TimeStamp:time.Now(),
		Status:1,
	}
	projson := cat.ToJson()
	if projson == "" {
		t.Errorf("Failed on converting  to JSON")

	}
	fmt.Println("This is the resulting JSON",projson)
}