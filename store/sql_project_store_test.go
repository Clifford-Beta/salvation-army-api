package store

import (
	"salv_prj/model"
	"testing"
	"time"
)

var projeSTore = SqlProjectStore{Database}

func TestSqlProjectStore_Create(t *testing.T) {
	proj := model.Project{
		School:      1,
		Name:        "Road",
		Description: "Roads to all hostels in the school",
		Start:       time.Now(),
		Duration:    100,
		Progress:    0,
		Status:      1,
		TimeStamp:   time.Now(),
	}
	res := <-projeSTore.Create(&proj)
	if res.Err != nil {
		t.Errorf("Project creation failed with", res.Err)
	}

}

func TestSqlProjectStore_Retrieve(t *testing.T) {
	res := <-projeSTore.Retrieve(1)
	if res.Err != nil {
		t.Errorf("Project retrieval failed with", res.Err)
	}
}

func TestSqlProjectStore_RetrieveAll(t *testing.T) {
	res := <-projeSTore.RetrieveAll()
	if res.Err != nil {
		t.Errorf("Projects retrieval failed with", res.Err)
	}
}
