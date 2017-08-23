package category

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
	Next           CategoryService
}

func (mw InstrumentingMiddleware) Create(category model.Category) (output *model.Category, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Create(category)
	return
}
func (mw InstrumentingMiddleware) CreateTier(category model.Tier) (output *model.Tier, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_tier", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.CreateTier(category)
	return
}

func (mw InstrumentingMiddleware) GetOne(id int) (output model.Category, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_one", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOne(id)
	return
}

func (mw InstrumentingMiddleware) GetOneTier(id int) (output model.Tier, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_one_tier", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOneTier(id)
	return
}

func (mw InstrumentingMiddleware) GetAll() (output []*model.Category, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_all", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAll()
	return
}
func (mw InstrumentingMiddleware) GetAllTiers() (output []*model.Tier, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_all_tiers", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAllTiers()
	return
}
