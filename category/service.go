package category

import (
	"salv_prj/model"
	"salv_prj/store"
)

type CategoryService interface {
	Create(category model.Category)(*model.Category,error)
	CreateTier(category model.Tier)(*model.Tier,error)
	GetOne(id int)(model.Category,error)
	GetOneTier(id int)(model.Tier,error)
	GetAll()([]*model.Category,error)
	GetAllTiers()([]*model.Tier,error)
}
type Categoryservice struct {}

func (Categoryservice)Create(category model.Category)(*model.Category,error)  {
	catStore := store.SqlCategoryStore{store.Database}
	category.Status = 1
	res := <-catStore.Save(&category)
	if res.Err != nil {
		return &model.Category{},res.Err
	}
	return res.Data.(*model.Category),nil
}
func (Categoryservice)CreateTier(category model.Tier)(*model.Tier,error)  {
	catStore := store.SqlCategoryStore{store.Database}
	res := <-catStore.CreateTier(&category)
	if res.Err != nil {
		return &model.Tier{},res.Err
	}
	return res.Data.(*model.Tier),nil
}

func (Categoryservice)GetOne(id int)(model.Category,error)  {
	catStore := store.SqlCategoryStore{store.Database}
	sch := <- catStore.Get(id)
	if sch.Err != nil {
		return model.Category{},sch.Err
	}
	return sch.Data.(model.Category),nil

}
func (Categoryservice)GetOneTier(id int)(model.Tier,error)  {
	catStore := store.SqlCategoryStore{store.Database}
	sch := <- catStore.GetTier(id)
	if sch.Err != nil {
		return model.Tier{},sch.Err
	}
	return sch.Data.(model.Tier),nil

}

func (Categoryservice) GetAll()([]*model.Category,error) {
	catStore := store.SqlCategoryStore{store.Database}
	sch := <- catStore.GetMany()
	if sch.Err != nil {
		return []*model.Category{},sch.Err
	}
	return sch.Data.([]*model.Category),nil

}

func (Categoryservice) GetAllTiers()([]*model.Tier,error) {
	catStore := store.SqlCategoryStore{store.Database}
	sch := <- catStore.GetManyTiers()
	if sch.Err != nil {
		return []*model.Tier{},sch.Err
	}
	return sch.Data.([]*model.Tier),nil

}