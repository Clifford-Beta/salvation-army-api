package store

import (
	"salvation-army-api/model"
	"testing"
	"time"
)

var fileStore = SqlFileStore{Database}

func TestSqlFileStore_CreateFileType(t *testing.T) {
	fltp := model.FileType{
		Name:        "Memo",
		Description: "These are memos",
		Store:       "Memoir",
		TimeStamp:   time.Now(),
	}
	res := <-fileStore.CreateFileType(&fltp)
	if res.Err != nil {
		t.Errorf("Create File Type Failed with", res.Err)
	}
}

func TestSqlFileStore_Create(t *testing.T) {
	inf := model.File{
		Name:        "Computer",
		Type:        1,
		Description: "Desktop computers",
		DateCreated: time.Now(),
		TimeStamp:   time.Now(),
	}
	res := <-fileStore.Create(&inf)
	if res.Err != nil {
		t.Errorf("Create File Failed with", res.Err)
	}

}

func TestSqlFileStore_RetrieveOne(t *testing.T) {
	res := <-fileStore.RetrieveOne(1)
	if res.Err != nil {
		t.Errorf("Retrieve File Failed with", res.Err)
	}
}

func TestSqlFileStore_RetrieveOneType(t *testing.T) {
	res := <-fileStore.RetrieveOneType(1)
	if res.Err != nil {
		t.Errorf("Retrieve File Type Failed with", res.Err)
	}
}

func TestSqlFileStore_RetrieveAll(t *testing.T) {
	res := <-fileStore.RetrieveAll()
	if res.Err != nil {
		t.Errorf("RetrieveAll Failed with", res.Err)

	}
}

func TestSqlFileStore_RetrieveAllTypes(t *testing.T) {
	res := <-fileStore.RetrieveAllTypes()
	if res.Err != nil {
		t.Errorf("RetrieveAllTypes Failed with", res.Err)

	}
}
