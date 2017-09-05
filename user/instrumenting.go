package user

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
	Next           UserService
}

func (mw InstrumentingMiddleware) Create(user model.User) (output *model.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Create(user)
	return
}

func (mw InstrumentingMiddleware) Update(user model.User) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Update(user)
	return
}

func (mw InstrumentingMiddleware) GetOne(id int) (output model.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getone", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOne(id)
	return
}

func (mw InstrumentingMiddleware) Login(email, password string) (output model.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "login", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Login(email, password)
	return
}

func (mw InstrumentingMiddleware) GetAll() (output map[string][]*model.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getall", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAll()
	return
}
