package file

import (
	"salv_prj/model"
	"salv_prj/store"
	"time"
)

// StringService provides operations on strings.
type FileService interface {
	Create(file model.File) (*model.File, error)
	CreateType(file model.FileType) (*model.FileType, error)
	GetOne(int) (model.File, error)
	GetOneType(int) (model.FileType, error)
	GetAll() (map[string][]*model.File, error)
	GetAllTypes() (map[string][]*model.FileType, error)
}

type Fileservice struct{}

func (Fileservice) Create(file model.File) (*model.File, error) {
	fileStore := store.SqlFileStore{store.Database}
	file.Status = 1
	file.DateCreated = time.Now()
	me := <- fileStore.Create(&file)
	if me.Err != nil {
		return &model.File{},me.Err
	}
	return me.Data.(*model.File), nil
}

func (Fileservice) CreateType(file model.FileType) (*model.FileType, error) {
	fileStore := store.SqlFileStore{store.Database}
	file.Status = 1
	me := <- fileStore.CreateFileType(&file)
	if me.Err != nil {
		return &model.FileType{},me.Err
	}
	return me.Data.(*model.FileType), nil
}




func (Fileservice) GetOneType( id int) (model.FileType, error) {
	fileStore := store.SqlFileStore{store.Database}
	me := <- fileStore.RetrieveOneType(id)
	if me.Err != nil {
		return model.FileType{},me.Err
	}
	return me.Data.(model.FileType), nil
}


func (Fileservice) GetOne(id int) (model.File, error) {
	fileStore := store.SqlFileStore{store.Database}
	me := <- fileStore.RetrieveOne(id)
	if me.Err != nil {
		return model.File{},me.Err
	}
	return me.Data.(model.File), nil
}


func (Fileservice) GetAll() (map[string][]*model.File, error) {
	fileStore := store.SqlFileStore{store.Database}
	me := <-fileStore.RetrieveAll()
	if me.Err != nil {
		return map[string][]*model.File{"data":[]*model.File{}},me.Err
	}
	return  map[string][]*model.File{"data":me.Data.([]*model.File)},nil
}

func (Fileservice) GetAllTypes() (map[string][]*model.FileType, error) {
	fileStore := store.SqlFileStore{store.Database}
	me := <-fileStore.RetrieveAllTypes()
	if me.Err != nil {
		return map[string][]*model.FileType{"data":[]*model.FileType{}},me.Err
	}
	return  map[string][]*model.FileType{"data":me.Data.([]*model.FileType)},nil
}





