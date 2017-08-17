package category

import (
	"salv_prj/model"
	"salv_prj/store"
	"time"
)

type CategoryServcie interface {
	Create(category model.Category)(*model.Category,error)
	GetOne(id int)(model.Category,error)
	GetAll()([]*model.Category,error)
}
type Categoryservice struct {}

func (Categoryservice)Create(category model.Category)(*model.Category,error)  {
	catStore := store.SqlCategoryStore{store.Database}
}