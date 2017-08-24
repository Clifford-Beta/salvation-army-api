package store

import "salvation-army-api/model"

type SqlInfrastructureStore struct {
	*SqlStore
}

func (s SqlInfrastructureStore) Create(inf *model.Infrastructure) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(inf); err != nil {

			result.Err = model.NewLocAppError("SqlInfrastructureStore.Create", "store.sql_inf.save.app_error", nil, err.Error())

		} else {
			//user.Sanitize()
			result.Data = inf
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlInfrastructureStore) CreateIType(inf *model.InfrastructureType) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(inf); err != nil {

			result.Err = model.NewLocAppError("SqlInfrastructureStore.CreateType", "store.sql_inf_type.save.app_error", nil, err.Error())

		} else {
			//user.Sanitize()
			result.Data = inf
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlInfrastructureStore) RetrieveOne(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var inf model.InfrastructureResult
		err := s.master.SelectOne(&inf, `select infrastructure.infrastructure_id as id, school.school_name as school,
												infrastructure_name as name, i_type.i_type_name as type,
												infrastructure_quantity as quantity, infrastructure_description as description,
												infrastructure.date_created as date_created
												from infrastructure
												inner join school on infrastructure.school_id = school.school_id
												inner join i_type on infrastructure.infrastructure_type = i_type.i_type_id
												where infrastructure.infrastructure_id = ? and infrastructure_status=1`, id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlInfrastructureStore.Get", "store.sql_inf.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = inf
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlInfrastructureStore) RetrieveOneType(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var inf model.InfrastructureType
		err := s.master.SelectOne(&inf, "select * from i_type where i_type_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlInfrastructureStore.GetType", "store.sql_inf_type.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = inf
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlInfrastructureStore) RetrieveAll() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var infs []model.InfrastructureResult
		_, err := s.master.Select(&infs, `select infrastructure.infrastructure_id as id, school.school_name as school,
			infrastructure_name as name, i_type.i_type_name as type,
			infrastructure_quantity as quantity, infrastructure_description as description,
				infrastructure.date_created as date_created
			from infrastructure
			inner join school on infrastructure.school_id = school.school_id
			inner join i_type on infrastructure.infrastructure_type = i_type.i_type_id
			where infrastructure_status=?`, 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlInfrastructureStore.GetAll", "store.sql_inf.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = infs
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlInfrastructureStore) RetrieveAllTypes() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var infs []model.InfrastructureType
		_, err := s.master.Select(&infs, "select * from i_type where i_type_status=?", 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlInfrastructureStore.GetAllTypes", "store.sql_inf_type.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		} else {
			if len(infs) == 0 {
				result.Err = model.NewLocAppError("SqlInfrastructureStore.RetrieveAll", "store.sql_infrastructure .retrieve_all.app_error", nil, "No records found")

			}
		}
		result.Data = infs
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}
