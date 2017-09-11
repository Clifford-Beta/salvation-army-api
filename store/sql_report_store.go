package store

import (
	"time"
	"salvation-army-api/model"
)
type SqlReportStore struct {
	*SqlStore
}

//dashboard data
type StatusResponse struct {
	Id int `json:"id" db:"id"`
	Status string `db:"status"`
	Count int `db:"count"`
}

type DashCountResponse struct {
	School int `json:"school" db:"school"`
	User int `db:"user" json:"user"`
	Message int `db:"msg" json:"message"`
}

type DashSchoolTrendResponse struct {
	Id int `db:id json:"id"`
	School string `json:"school" db:"school"`
	Year int `db:"year" json:"year"`
	Mark float64 `db:"mark" json:"mark"`
}

type SchoolPerformanceByCatResponse struct {
	Rank int `db:"rank" json:"rank"`
	Id int `db:id json:"id"`
	School string `json:"school" db:"school"`
	Category string `json:"category" db:"category"`
	Year int `db:"year" json:"year"`
	Mark float64 `db:"mark" json:"mark"`
}

func (s SqlReportStore) ProjectStatus(from,to time.Time) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var school []*StatusResponse
		_,err := s.GetMaster().Select(&school, `
							SELECT sub.*,count(*) over(partition by sub.status) as count
							FROM
							(SELECT project_id as id,
							CASE
							WHEN DATE_ADD(project_start, INTERVAL project_duration DAY) >= NOW() THEN "green"
							ELSE "red"
							END
							"status"
							FROM project
							INNER JOIN school on project.school = school.school_id
							WHERE project.timestamp BETWEEN :From AND :To
							AND project_status = 1
							) sub`, map[string]interface{}{"From":from,"To":to,})
		if err != nil  {
			result.Err = model.NewLocAppError("SqlReportStore.ProjectStatus", "store.sql_school.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = school

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}


func (s SqlReportStore) DashCount() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var res DashCountResponse
		err := s.GetMaster().SelectOne(&res, `
							select count(distinct user.user_id) as user,
							   count(distinct message.message_id) as msg,
							   count(distinct school.school_id) as school
								from user,message, school
							`)
		if err != nil  {
			result.Err = model.NewLocAppError("SqlReportStore.DashCount", "store.sql_school.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = res
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s SqlReportStore) DashSchoolTrend() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var res []*DashSchoolTrendResponse
		_,err := s.GetMaster().Select(&res, `
							select s_performance_id as id,school_name as school ,max(s_performance_mark)
							over (partition by s_performance_year ) as mark,s_performance_year as year from s_performance
							left join school on school=school_id
							where s_performance_status =1
							order by s_performance_year asc
							`)
		if err != nil  {
			result.Err = model.NewLocAppError("SqlReportStore.DashSchoolTrend", "store.sql_school.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = res
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}


func (s SqlReportStore) SchoolRankingByCategory(category int,from,to time.Time) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var res []*SchoolPerformanceByCatResponse
		filters := map[string]interface{}{}
		filters["from"]=from
		filters["to"]=to
		filters["category"] = category

		_,err := s.GetMaster().Select(&res, `
							select * from (
						select rank() over (partition by category_name order by s_performance.s_performance_mark desc) as rank, school.school_id as id, school.school_name as school ,s_performance.s_performance_mark as mark,
						s_performance_year as year, category_name as category
						from school
						inner join s_performance on school = school_id
						left join category on s_performance_cat = category_id
						inner join tier on school_category = tier_id
						where school_status = 1  AND date_registered BETWEEN :from AND :to
						AND s_performance_cat = :category
						) as school_ranks
						order by rank
							`,filters)
		if err != nil  {
			result.Err = model.NewLocAppError("SqlReportStore.SchoolRankingByCategory", "store.sql_school.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = res
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

