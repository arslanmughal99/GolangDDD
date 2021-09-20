package shared

import (
	"sync"

	"massivleads/domain/services/shared/session/redis"
	"massivleads/prototypes/services/shared/session"
)

type Service struct {
	Session session.Service
}

var (
	services   = new(Service)
	sharedOnce = sync.Once{}
)

func GetSharedServices() Service {
	sharedOnce.Do(
		func() {
			services.Session = redis.NewRedisSessionService()
		},
	)

	return *services
}
