package file


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
	Next           FileService
}

func (mw InstrumentingMiddleware) Create(file model.File) (output *model.File, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Create(file)
	return
}

func (mw InstrumentingMiddleware) CreateType(file model.FileType) (output *model.FileType, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_type", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.CreateType(file)
	return
}

func (mw InstrumentingMiddleware) GetOne( id int) (output model.File, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_one", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOne(id)
	return
}

func (mw InstrumentingMiddleware) GetOneType( id int) (output model.FileType, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_one", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetOneType(id)
	return
}



func (mw InstrumentingMiddleware) GetAll() (output map[string][]*model.File, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getall", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAll()
	return
}

func (mw InstrumentingMiddleware) GetAllTypes() (output map[string][]*model.FileType, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_all", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetAllTypes()
	return
}