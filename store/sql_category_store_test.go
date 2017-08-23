package store

import (
	"salv_prj/model"
	"testing"
)

var catStore = SqlCategoryStore{Database}

func TestSqlCategoryStore_Save(t *testing.T) {
	cat := model.Category{Name: "Mine", Description: "This is a test category"}
	res := <-catStore.Save(&cat)
	if res.Err != nil {
		t.Errorf("category save test failed with error ", res.Err)
	}
}
func TestSqlCategoryStore_Get(t *testing.T) {
	res := <-catStore.Get(1)
	if res.Err != nil {
		t.Errorf("category get test failed with error ", res.Err)
	}
}

func TestSqlCategoryStore_GetMany(t *testing.T) {
	res := <-catStore.GetMany()
	if res.Err != nil {
		t.Errorf("category get many test failed with error ", res.Err)
	}

}
