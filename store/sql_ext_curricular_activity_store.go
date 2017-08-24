package store

import (
	"salvation-army-api/model"
)

type SqlExtraCurricularStore struct {
	*SqlStore
}

func (s SqlExtraCurricularStore) CreateNewActivity(act *model.ExtraCurricular) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(act); err != nil {
			result.Err = model.NewLocAppError("SqlExtraCurricularStore.Save", "store.sql_extra_curricular.save.app_error", nil, "ext="+act.Name+", "+err.Error())

		} else {
			result.Data = act
		}
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel

}

func (s SqlExtraCurricularStore) GetActivity(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var act model.ExtraCurricular
		err := s.GetMaster().SelectOne(&act, "select * from ext_curricular where ext_curricular_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlExtraCurricularStore.Get", "store.sql_extracurricular.get.app_error", nil, "extcurricular="+act.Name+", "+err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		//user.Sanitize()
		result.Data = act

		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlExtraCurricularStore) GetAllActivities() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var acts []*model.ExtraCurricular
		_, err := s.GetMaster().Select(&acts, "select * from ext_curricular where ext_curricular_status=?", 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlExtraCurricularStore.GetAll", "store.sql_extracurricular.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		//user.Sanitize()
		result.Data = acts

		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlExtraCurricularStore) GetRecordedActivity(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var act model.ExtraCurricularActivity
		err := s.GetMaster().SelectOne(&act, "select * from ext_activity where ext_activity_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlExtraCurricularActivityStore.Get", "store.sql_extracurricular.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		//user.Sanitize()
		result.Data = act

		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlExtraCurricularStore) GetAllRecordedActivities() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var acts []*model.ExtraCurricularActivityResult
		_, err := s.GetMaster().Select(&acts, `select ext_curricular_id as id, ext_curricular_name as name,ext_curricular_desc as description,
													ext_activity_performance, date, ext_level_name as level, ext_level_desc as level_description , school.school_name
													FROM ext_activity
													inner join ext_curricular on ext_activity.activity = ext_curricular.ext_curricular_id
													inner join ext_level on ext_level.ext_level_id = ext_activity.level
													inner join school on ext_activity.school=school.school_id
													where ext_activity_status = ?`, 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlExtraCurricularActivityStore.GetAll", "store.sql_extracurricular.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		result.Data = acts

		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlExtraCurricularStore) RecordActivity(activity *model.ExtraCurricularActivity) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(activity); err != nil {
			result.Err = model.NewLocAppError("SqlExtraCurricularStore.RecordActivity", "store.sql_extra_curricular.save.app_error", nil, err.Error())

		} else {
			result.Data = activity
		}
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlExtraCurricularStore) RecordLevel(level *model.ExtraCurricularLevel) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(level); err != nil {
			result.Err = model.NewLocAppError("SqlExtraCurricularStore.RecordLevel", "store.sql_extra_curricular.save.app_error", nil, err.Error())

		} else {
			result.Data = level
		}
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlExtraCurricularStore) GetLevel(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var act model.ExtraCurricularLevel
		err := s.GetMaster().SelectOne(&act, "select * from ext_level where ext_level_id=?", id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlExtraCurricularLevelStore.Get", "store.sql_extracurricular.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}

		//pl.AddUser(&user)
		//user.Sanitize()
		result.Data = act

		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

func (s SqlExtraCurricularStore) GetAllLevels() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var acts []*model.ExtraCurricularLevel
		_, err := s.GetMaster().Select(&acts, "select * from ext_level where ext_level_status=?", 1)
		if err != nil {
			result.Err = model.NewLocAppError("SqlExtraCurricularLevelStore.GetAll", "store.sql_extracurricular.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		} else {
			if len(acts) == 0 {
				result.Err = model.NewLocAppError("SqlBestTeacherStore.GetMany", "store.sql_best_teacher.getmany.app_error", nil, "No records found")

			}
		}
		result.Data = acts

		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}
