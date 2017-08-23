package store

import (
	"salv_prj/model"
	"testing"
	"time"
)

var strStore = SqlExtraCurricularStore{Database}

func TestSqlExtraCurricularStore_CreateNewActivity(t *testing.T) {
	extcr := model.ExtraCurricular{
		Id:          1,
		Name:        "Drama",
		Description: "Acting and drama festivals",
		TimeStamp:   time.Now(),
		Status:      1,
	}
	res := <-strStore.CreateNewActivity(&extcr)
	if res.Err != nil {
		t.Errorf("Create New Extra Curricular activity failed with ", res.Err)
	}
}

func TestSqlExtraCurricularStore_GetActivity(t *testing.T) {
	res := <-strStore.GetActivity(1)
	if res.Err != nil {
		t.Errorf("ExtraCurricularStore_GetActivity failed with", res.Err)
	}
}

func TestSqlExtraCurricularStore_GetAllActivities(t *testing.T) {
	res := <-strStore.GetAllActivities()
	if res.Err != nil {
		t.Errorf("ExtraCurricularStore_GetAllActivities failed with", res.Err)
	}

}

func TestSqlExtraCurricularStore_RecordLevel(t *testing.T) {
	extlvl := model.ExtraCurricularLevel{
		Id:          1,
		Name:        "Nationals",
		Description: "Nationwide competitions",
		TimeStamp:   time.Now(),
		Status:      1,
	}
	res := <-strStore.RecordLevel(&extlvl)
	if res.Err != nil {
		t.Errorf("ExtraCurricularStore_RecordLevel failed with", res.Err)
	}
}

func TestSqlExtraCurricularStore_GetLevel(t *testing.T) {
	res := <-strStore.GetLevel(1)
	if res.Err != nil {
		t.Errorf("ExtraCurricularStore_GetLevel failed with", res.Err)
	}
}

func TestSqlExtraCurricularStore_GetAllLevels(t *testing.T) {
	res := <-strStore.GetAllLevels()
	if res.Err != nil {
		t.Errorf("ExtraCurricularStore_GetAllLevels failed with", res.Err)
	}
}

func TestSqlExtraCurricularStore_RecordActivity(t *testing.T) {
	extact := model.ExtraCurricularActivity{
		Id:          1,
		School:      1,
		Level:       1,
		Activity:    1,
		Performance: "First Runners Up",
		Date:        time.Now(),
		TimeStamp:   time.Now(),
		Status:      1,
	}

	res := <-strStore.RecordActivity(&extact)
	if res.Err != nil {
		t.Errorf("ExtraCurricularStore_RecordActivity failed with", res.Err)
	}
}

func TestSqlExtraCurricularStore_GetRecordedActivity(t *testing.T) {
	res := <-strStore.GetRecordedActivity(1)
	if res.Err != nil {
		t.Errorf("ExtraCurricularStore_GetRecordedActivity failed with", res.Err)
	}
}

func TestSqlExtraCurricularStore_GetAllRecordedActivities(t *testing.T) {
	res := <-strStore.GetAllRecordedActivities()
	if res.Err != nil {
		t.Errorf("ExtraCurricularStore_GetAllRecordedActivities failed with", res.Err)
	}
}
