package message

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
	Next           MessageService
}

func (mw InstrumentingMiddleware) Create(message model.Message) (output *model.Message, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Create(message)
	return
}

func (mw InstrumentingMiddleware) Update(message model.Message) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Update(message)
	return
}

func (mw InstrumentingMiddleware) GetOne(id int) (output model.Message, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getone", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOne(id)
	return
}

func (mw InstrumentingMiddleware) GetAll(user int) (output map[string][]model.MessageResult, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getall", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAll(user)
	return
}
