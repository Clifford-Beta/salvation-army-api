package store

import (
	"salvation-army-api/model"
	"testing"
	"time"
)

var infStore = SqlInfrastructureStore{Database}

func TestSqlInfrastructureStore_Create(t *testing.T) {
	inf := model.Infrastructure{
		School:      1,
		Name:        "Computer",
		Type:        1,
		Quantity:    100,
		Description: "Desktop computers",
		DateCreated: time.Now(),
		TimeStamp:   time.Now(),
	}
	res := <-infStore.Create(&inf)
	if res.Err != nil {
		t.Errorf("Create Infrastructure Failed with", res.Err)
	}

}

func TestSqlInfrastructureStore_CreateIType(t *testing.T) {
	inf := model.InfrastructureType{
		Name:        "Electronics",
		Description: "Electical appliances",
		TimeStamp:   time.Now(),
	}
	res := <-infStore.CreateIType(&inf)
	if res.Err != nil {
		t.Errorf("Create Infrastructure Type Failed with", res.Err)
	}

}

func TestSqlInfrastructureStore_RetrieveOne(t *testing.T) {
	res := <-infStore.RetrieveOne(1)
	if res.Err != nil {
		t.Errorf("Retrieve Failed with", res.Err)
	}
}

func TestSqlInfrastructureStore_RetrieveOneType(t *testing.T) {
	res := <-infStore.RetrieveOneType(1)
	if res.Err != nil {
		t.Errorf("Retrieve Type Failed with", res.Err)
	}
}

func TestSqlInfrastructureStore_RetrieveAll(t *testing.T) {
	res := <-infStore.RetrieveAll()
	if res.Err != nil {
		t.Errorf("RetrieveAll Failed with", res.Err)

	}
}

func TestSqlInfrastructureStore_RetrieveAllTypes(t *testing.T) {
	res := <-infStore.RetrieveAllTypes()
	if res.Err != nil {
		t.Errorf("RetrieveAllTypes Failed with", res.Err)

	}
}
