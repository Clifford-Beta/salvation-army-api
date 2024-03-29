package school

import (
	"fmt"
	"github.com/go-kit/kit/metrics"
	"salvation-army-api/model"
	"time"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           SchoolService
}

func (mw InstrumentingMiddleware) Create(school model.School) (output *model.School, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Create(school)
	return
}

func (mw InstrumentingMiddleware) Update(school model.School) (output UpdateResponse, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Update(school)
	return
}

func (mw InstrumentingMiddleware) Delete(school model.School) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Delete(school)
	return
}

func (mw InstrumentingMiddleware) RecordPerformance(performance *model.SchoolPerformance) (output *model.SchoolPerformance, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "RecordPerformance", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.RecordPerformance(performance)
	return
}
func (mw InstrumentingMiddleware) GetBestSchool(from, to int) (output model.SchoolPerformanceResult, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetBestSchool", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetBestSchool(from, to)
	return
}

func (mw InstrumentingMiddleware) GetOne(id int) (output model.School, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getone", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOne(id)
	return
}

//func (mw InstrumentingMiddleware) GetAll() (output []*model.School, err error) {
//	defer func(begin time.Time) {
//		lvs := []string{"method", "getall", "error", fmt.Sprint(err != nil)}
//		mw.RequestCount.With(lvs...).Add(1)
//		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
//	}(time.Now())
//
//	output, err = mw.Next.GetAll()
//	return
//}
func (mw InstrumentingMiddleware) GetAll() (output map[string][]*model.SchoolResult, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getall", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAll()
	return
}

func (mw InstrumentingMiddleware) RankAllSchools(from, to int) (output map[string][]model.SchoolPerformanceResult, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "RankAllSchools", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.RankAllSchools(from, to)
	return
}

func (mw InstrumentingMiddleware) GetDashboardData() (output map[string]interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetDashboardData", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetDashboardData()
	return
}
func (mw InstrumentingMiddleware) RankSchoolByCategory(category int, from,to time.Time) (output map[string]interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "RankSchoolByCategory", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.RankSchoolByCategory(category,from,to)
	return
}
