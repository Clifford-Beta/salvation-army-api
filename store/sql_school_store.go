package store

import (
	"net/http"
	"salvation-army-api/model"
	"strconv"
)

type SqlSchoolStore struct {
	*SqlStore
}

func (s SqlSchoolStore) Save(school *model.School) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(school); err != nil {
			if IsUniqueConstraintError(err.Error(), []string{"Email", "school_email_key", "idx_users_email_unique"}) {
				result.Err = model.NewAppError("SqlSchoolStore.Save", "store.sql_school.save.email_exists.app_error", nil, err.Error(), http.StatusBadRequest)
			} else if IsUniqueConstraintError(err.Error(), []string{"Phone", "school__phone_key", "idx_school_phone_unique"}) {
				result.Err = model.NewAppError("SqlSchoolStore.Save", "store.sql__school.save.phone_exists.app_error", nil, err.Error(), http.StatusBadRequest)
			} else {
				result.Err = model.NewLocAppError("SqlSchoolStore.Save", "store.sql_school.save.app_error", nil, "school="+school.Name+", "+err.Error())
			}
		} else {
			result.Data = school
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlSchoolStore) RecordPerformance(p *model.SchoolPerformance) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if err := s.GetMaster().Insert(p); err != nil {

			result.Err = model.NewLocAppError("SqlSchoolStore.RecordPerformance", "store.sql_school.save.app_error", nil, err.Error())

		} else {
			result.Data = p
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlSchoolStore) RetrieveBestPerfomingSchool(filter map[string]interface{}) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	var err error
	go func() {
		result := StoreResult{}
		var sch model.SchoolPerformanceResult
		if _, ok := filter["From"]; ok {
			if _, ok := filter["To"]; ok {
				err = s.master.SelectOne(&sch,
					`select school.school_id as id, school.school_name as school,school.school_location as location,
						school.school_description as description,  tier_name as tier ,max(s_performance.s_performance_mark) as mark,
						s_performance_year as year, category_name as category,
						school.date_registered
						from school
						inner join s_performance on school = school_id
						left join category on s_performance_cat = category_id
						inner join tier on school_category = tier_id
						where school_status = 1  and s_performance_year between :From and :To `, filter)
			} else {
				result.Err = model.NewLocAppError("SqlSqlSchoolStoreStore.Get", "The filter value for year is not a valid year", nil, err.Error())
				storeChannel <- result
				close(storeChannel)
				return

			}
		} else {
			err = s.master.SelectOne(&sch,
				`select school.school_id as id, school.school_name as school,school.school_location as location,
						school.school_description as description,  tier_name as tier ,max(s_performance.s_performance_mark) as mark,
						s_performance_year as year, category_name as category,
						school.date_registered
						from school
						inner join s_performance on school = school_id
						left join category on s_performance_cat = category_id
						inner join tier on school_category = tier_id
						where school_status = 1 `)
		}
		if err != nil {
			result.Err = model.NewLocAppError("SqlSqlSchoolStoreStore.Get", "store.sql_school.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = sch
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}



func (s SqlSchoolStore) Update(school *model.School) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		if sqlResult, err := s.GetMaster().Exec(
			`UPDATE
				school
			SET
				school_name = :Name,
				school_email = :Email,
				school_phone = :Phone,
				school_postal_address = :PostalAddress,
				school_category = :Category,
				school_logo = :Logo,
				school_location = :Location,
				school_description = :Description
			WHERE
				school_id = :Id`,
			map[string]interface{}{
				"Id":             school.Id,
				"Name":      school.Name,
				"Email": school.Email,
				"Phone":         school.Phone,
				"PostalAddress":           school.PostalAddress,
				"Category":       school.Category,
				"Logo":       school.Logo,
				"Location":       school.Location,
				"Description":       school.Description,
			}); err != nil {
			result.Err = model.NewLocAppError("SqlSchoolStore.UpdateOptimistically",
				"store.sql_job.update.app_error", nil, "id="+strconv.Itoa(school.Id)+", "+err.Error())
		} else {
			rows, err := sqlResult.RowsAffected()

			if err != nil {
				result.Err = model.NewLocAppError("SqlSchoolStore.UpdateOptimistically",
					"store.sql_job.update.app_error", nil, "id="+strconv.Itoa(school.Id)+", "+err.Error())
			} else {
				if rows == 1 {
					result.Data = true
				} else {
					result.Data = false
				}
			}
		}

		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}



func (s SqlSchoolStore) Delete(school *model.School) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		res, err := s.GetMaster().Exec("Update school SET school_status=0 where school_id=?", school.Id)
		if err != nil {
			result.Err = model.NewLocAppError("SqlSchoolStore.Delete", "store.sql_school.delete.app_error", nil, "user_id="+strconv.Itoa(school.Id)+", "+err.Error())

		} else {
			result.Data = res
		}
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (s SqlSchoolStore) Get(id int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var school model.School
		err := s.master.SelectOne(&school, "select * from school where school_id=?", id)
		if err != nil  {
			result.Err = model.NewLocAppError("SqlSchoolStore.Get", "store.sql_school.get.app_error", nil, "school="+school.Name+", "+err.Error())
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

func (s SqlSchoolStore) GetMany() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var schools []*model.SchoolResult
		_, err := s.GetMaster().Select(&schools, `select school_id as id,school_name as name, school_postal_address as postal_address, school_phone as phone,
														school_logo as logo, school_email as email, school_location as location,
															school_description as description,date_registered,tier_name as category
															from school
															inner join tier on school.school_category = tier_id
															where school_status = 1`)
		if err != nil {
			result.Err = model.NewLocAppError("SqlUsertore.GetMany", "store.sql_school.getmany.app_error", nil, err.Error())

		} else {
			if len(schools) == 0 {
				result.Err = model.NewLocAppError("SqlSchoolStore.GetMany", "store.sql_school.get_many.app_error", nil, "No records found")

			}
			result.Data = schools
		}
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}


//ranking
//rank all schools
func (s SqlSchoolStore) RankAllSchools() StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var sch []model.SchoolPerformanceResult
		_, err := s.master.Select(&sch, `select * from (
											select rank() over (partition by category_name order by s_performance.s_performance_mark desc) as rank, school.school_id as id, school.school_name as school,school.school_location as location,
											school.school_description as description,  tier_name as tier ,s_performance.s_performance_mark as mark,
											s_performance_year as year, category_name as category,
											school.date_registered
											from school
											inner join s_performance on school = school_id
											left join category on s_performance_cat = category_id
											inner join tier on school_category = tier_id
											where school_status = 1 ) as school_ranks
											order by rank`)

		if err != nil {
			result.Err = model.NewLocAppError("SqlSqlSchoolStoreStore.RankAllSchools", "store.sql_school.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = sch
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

//rank all schools for a year

func (s SqlSchoolStore) RankAllSchoolsPerYear(year int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var sch []model.SchoolPerformanceResult
		_, err := s.master.Select(&sch, `select * from (
											select rank() over (partition by category_name order by s_performance.s_performance_mark desc) as rank, school.school_id as id, school.school_name as school,school.school_location as location,
											school.school_description as description,  tier_name as tier ,s_performance.s_performance_mark as mark,
											s_performance_year as year, category_name as category,
											school.date_registered
											from school
											inner join s_performance on school = school_id
											left join category on s_performance_cat = category_id
											inner join tier on school_category = tier_id
											where school_status = 1 ) as school_ranks
											where (school_ranks.year = :year)
											order by rank`,map[string]interface{}{"year":year})
		if err != nil {
			result.Err = model.NewLocAppError("SqlSqlSchoolStoreStore.RankAllSchoolsPerYear", "store.sql_school.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = sch
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

//rank schools based on category

func (s SqlSchoolStore) RankAllSchoolsPerCategory(tier int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var sch []model.SchoolPerformanceResult
		_, err := s.master.Select(&sch, `select * from (
							select rank() over (partition by category_name, tier_name order by s_performance.s_performance_mark desc) as rank, school.school_id as id, school.school_name as school,school.school_location as location,
							school.school_description as description,  tier_name as tier ,s_performance.s_performance_mark as mark,
							s_performance_year as year, category_name as category,
							school.date_registered
							from school
							inner join s_performance on school = school_id
							left join category on s_performance_cat = category_id
							inner join tier on school_category = tier_id
							where school_status = 1 and school_category = :tier ) as school_ranks
							order by rank`,map[string]interface{}{"tier":tier})
		if err != nil {
			result.Err = model.NewLocAppError("SqlSqlSchoolStoreStore.RankAllSchoolsPerCategory", "store.sql_school.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = sch
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}

// category and year

func (s SqlSchoolStore) RankAllSchoolsPerCategoryAndYear(tier,from,to int) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		var sch []model.SchoolPerformanceResult
		_, err := s.master.Select(&sch, `select * from (
							select rank() over (partition by category_name, tier_name order by s_performance.s_performance_mark desc) as rank, school.school_id as id, school.school_name as school,school.school_location as location,
							school.school_description as description,  tier_name as tier ,s_performance.s_performance_mark as mark,
							s_performance_year as year, category_name as category,
							school.date_registered
							from school
							inner join s_performance on school = school_id
							left join category on s_performance_cat = category_id
							inner join tier on school_category = tier_id
							where school_status = 1 and school_category = :tier ) as school_ranks
							order by rank`,map[string]interface{}{"tier":tier})
		if err != nil {
			result.Err = model.NewLocAppError("SqlSqlSchoolStoreStore.RankAllSchoolsPerCategory", "store.sql_school.get.app_error", nil, err.Error())
			storeChannel <- result
			close(storeChannel)
			return
		}
		result.Data = sch
		storeChannel <- result
		close(storeChannel)

	}()
	return storeChannel
}
//get the top n schools in each tier

//select * from (
//select rank() over (partition by category_name, tier_name order by s_performance.s_performance_mark desc) as rank, school.school_id as id, school.school_name as school,school.school_location as location,
//school.school_description as description,  tier_name as tier ,s_performance.s_performance_mark as mark,
//s_performance_year as year, category_name as category,
//school.date_registered
//from school
//inner join s_performance on school = `school_id`
//left join category on s_performance_cat = category_id
//inner join tier on school_category = tier_id
//where school_status = 1 ) as school_ranks
//where (school_ranks.rank <= 3)
//order by tier, rank

////get top n schools overally
//
//select * from (
//select rank() over (partition by category_name, order by s_performance.s_performance_mark desc) as rank, school.school_id as id, school.school_name as school,school.school_location as location,
//school.school_description as description,  tier_name as tier ,s_performance.s_performance_mark as mark,
//s_performance_year as year, category_name as category,
//school.date_registered
//from school
//inner join s_performance on school = `school_id`
//left join category on s_performance_cat = category_id
//inner join tier on school_category = tier_id
//where school_status = 1 ) as school_ranks
//where (school_ranks.rank <= 3)
//order by rank