package auth

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type LoggingAuthMiddleware struct {
	Logger log.Logger
	Next   Authservice
}

func (mw LoggingAuthMiddleware) Auth(clientID string, clientSecret string) (token map[string]interface{}, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"clientID": clientID,
			"token":    token,
			"err":      err,
			"took":     time.Since(begin)}).Info("service = ", "auth ", "method = ", "auth")
	}(time.Now())

	token, err = mw.Next.Auth(clientID, clientSecret)
	return
}
