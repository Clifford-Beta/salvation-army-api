package store

import (
	"salv_prj/model"
	"strconv"
)

type SqlCategoryStore struct {
	*SqlStore
}

func (s SqlCategoryStore) Save(category *model.Category) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(category); err != nil {

			result.Err = model.NewLocAppError("SqlCategoryStore.Save", "store.sql_category.save.app_error", nil, "category="+category.Name+", "+err.Error())

		} else {
			//user.Sanitize()
			result.Data = category
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlCategoryStore) Delete(category *model.Category) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Exec("Update category SET category_status=0 where category_id=?", category.Id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlCategoryStore.Delete", "store.sql_category.delete.app_error", nil, "category_id="+strconv.Itoa(category.Id)+", "+err.Error())

		} else {
			result.Data = res
			//result.Err =
		}
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlCategoryStore) Get(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var category model.Category
		err := s.master.SelectOne(&category, "select * from category where category_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlCategoryStore.Get", "store.sql_category.get.app_error", nil, "category="+category.Name+", "+err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		//user.Sanitize()
		result.Data = category

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlCategoryStore) GetMany() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var categories []*model.Category
		_, err := s.GetMaster().Select(&categories, "SELECT * FROM category WHERE category_status=1")
		if err != nil {
			result.Err = model.NewLocAppError("SqlCategorytore.GetMany", "store.sql_category.getmany.app_error", nil, err.Error())

		} else {
			if len(categories) == 0 {
				result.Err = model.NewLocAppError("SqlCategorytore.GetMany", "store.sql_category.getmany.app_error", nil, "No records found")

			} else {
				result.Data = categories
			}

		}
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlCategoryStore) CreateTier(category *model.Tier) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(category); err != nil {

			result.Err = model.NewLocAppError("SqlCategoryStore.CreateTier", "store.sql_category.create_tier.app_error", nil, "category="+category.Name+", "+err.Error())

		} else {
			//user.Sanitize()
			result.Data = category
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlCategoryStore) DeleteTier(category *model.Tier) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Exec("Update tier SET tier_status=0 where tier_id=?", category.Id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlCategoryStore.DeleteTier", "store.sql_category.delete_tier.app_error", nil, "category_id="+strconv.Itoa(category.Id)+", "+err.Error())

		} else {
			result.Data = res
			//result.Err =
		}
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlCategoryStore) GetTier(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		//pl := model.NewUserList()
		var category model.Tier
		err := s.master.SelectOne(&category, "select * from tier where tier_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlCategoryStore.GetTier", "store.sql_category.get_tier.app_error", nil, "category="+category.Name+", "+err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		//user.Sanitize()
		result.Data = category

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlCategoryStore) GetManyTiers() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var categories []*model.Tier
		_, err := s.GetMaster().Select(&categories, "SELECT * FROM tier WHERE tier_status=1")
		if err != nil {
			result.Err = model.NewLocAppError("SqlCategorytore.GetManyTier", "store.sql_category.get_many_tiers.app_error", nil, err.Error())

		} else {
			if len(categories) == 0 {
				result.Err = model.NewLocAppError("SqlCategorytore.GetMany", "store.sql_category.getmany.app_error", nil, "No records found")

			}
			result.Data = categories

		}
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}
