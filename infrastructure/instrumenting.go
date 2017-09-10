package infrastructure

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
	Next           InfrastructureService
}

func (mw InstrumentingMiddleware) Create(inf model.Infrastructure) (output *model.Infrastructure, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Create(inf)
	return
}

func (mw InstrumentingMiddleware) Update(inf model.Infrastructure) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Update(inf)
	return
}

func (mw InstrumentingMiddleware) Delete(inf model.Infrastructure) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Delete(inf)
	return
}

func (mw InstrumentingMiddleware) CreateType(inf model.InfrastructureType) (output *model.InfrastructureType, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_type", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.CreateType(inf)
	return
}

func (mw InstrumentingMiddleware) GetOne(id int) (output model.Infrastructure, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_one", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOne(id)
	return
}

func (mw InstrumentingMiddleware) GetOneType(id int) (output model.InfrastructureType, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_one", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOneType(id)
	return
}

func (mw InstrumentingMiddleware) GetAll() (output map[string][]model.InfrastructureResult, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getall", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAll()
	return
}

func (mw InstrumentingMiddleware) GetAllTypes() (output map[string][]model.InfrastructureType, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_all", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAllTypes()
	return
}
