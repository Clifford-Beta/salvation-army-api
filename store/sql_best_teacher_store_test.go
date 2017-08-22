package store

import (
	"testing"
	"salv_prj/model"
)
var bestTStore = SqlBestTeacherStore{Database}

func TestSqlBestTeacherStore_Save(t *testing.T) {
	bst := model.BestTeacher{Name:"Mine",Category:2}
	res := <- bestTStore.Save(&bst)
	if res.Err != nil {
		t.Errorf("best teacher save test failed with error ",res.Err)
	}
}
func TestSqlBestTeacherStore_Get(t *testing.T) {
	res := <- bestTStore.Get(2015,2019)
	if res.Err != nil {
		t.Errorf("best teacher get test failed with error ",res.Err)
	}
}

func TestSqlBestTeacherStore_GetMany(t *testing.T) {
	res := <- bestTStore.GetMany()
	if res.Err != nil {
		t.Errorf("best teacher get many test failed with error ",res.Err)
	}

}