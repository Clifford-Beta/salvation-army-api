package auth

import (
	"fmt"
	"github.com/go-kit/kit/metrics"
	"time"
)

type InstrumentingAuthMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           AuthService
}

func (mw InstrumentingAuthMiddleware) Auth(clientID int, clientSecret string) (token string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "auth", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	token, err = mw.Next.Auth(clientID, clientSecret)
	return
}
