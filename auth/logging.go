package auth

import (
	"time"
	log "github.com/sirupsen/logrus"
)

type LoggingAuthMiddleware struct {
	Logger log.Logger
	Next   Authservice
}

func (mw LoggingAuthMiddleware) Auth(clientID string, clientSecret string) (token string, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"clientID": clientID,
			"token": token,
			"err": err,
			"took": time.Since(begin)}).Info( "service = ","auth ","method = ", "auth")
	}(time.Now())

	token, err = mw.Next.Auth(clientID, clientSecret)
	return
}