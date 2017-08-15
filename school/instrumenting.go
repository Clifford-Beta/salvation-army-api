package school

import (
"fmt"
"time"
"github.com/go-kit/kit/metrics"
"salv_prj/model"
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

func (mw InstrumentingMiddleware) GetOne(id int) (output model.School, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getone", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOne(id)
	return
}

func (mw InstrumentingMiddleware) GetAll() (output []*model.School, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getall", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAll()
	return
}