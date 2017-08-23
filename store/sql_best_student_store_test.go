package store

import (
	"salvation-army-api/model"
	"testing"
)

var bestStore = SqlBestStudentStore{Database}

func TestSqlBestStudentStore_Save(t *testing.T) {
	bst := model.BestStudent{Name: "Mine", Category: 2}
	res := <-bestStore.Save(&bst)
	if res.Err != nil {
		t.Errorf("best student save test failed with error ", res.Err)
	}
}
func TestSqlBestStudentStore_Get(t *testing.T) {
	res := <-bestStore.Get(2015, 2016)
	if res.Err != nil {
		t.Errorf("best student get test failed with error ", res.Err)
	}
}

func TestSqlBestStudentStore_GetMany(t *testing.T) {
	res := <-bestStore.GetMany()
	if res.Err != nil {
		t.Errorf("best student get many test failed with error ", res.Err)
	}

}
