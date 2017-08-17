package store

import "salv_prj/model"

type SqlProjectStore struct {
	*SqlStore
}

func (s SqlProjectStore)Create(project *model.Project) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(project); err != nil {

			result.Err = model.NewLocAppError("SqlProjectStore.Create", "store.sql_project.save.app_error", nil, err.Error())

		} else {
			result.Data = project
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}



func (s SqlProjectStore)Retrieve(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var project model.Project
		err := s.master.SelectOne(&project,"select * from project where project_id=?",id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlProjectStore.GetType", "store.sql_project_type.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = project
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlProjectStore)RetrieveAll()StoreChannel  {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var project []model.ProjectResult
		_,err := s.master.Select(&project,
			`
			select project.project_id as id, school.school_name as school, project.project_name as name,
			project.project_start as start, project.project_duration as duration,
			project.project_progress as progress, project.timestamp as time_stamp

			from project
			inner join school on school.school_id = project.school
			where project_status=1`)
		if err != nil {
			result.Err = model.NewLocAppError("SqlProjectStore.GetAll", "store.sql_project.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = project
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

