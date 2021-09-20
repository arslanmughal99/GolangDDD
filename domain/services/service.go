package services

import (
	"sync"

	"massivleads/domain/services/auth"
	"massivleads/domain/services/shared"
	"massivleads/domain/services/user"
	"massivleads/prototypes/services"
	sauth "massivleads/prototypes/services/auth"
)

type Services struct {
	User   services.User
	Auth   sauth.Service
	Shared shared.Service
}

var (
	srvs         = new(Services)
	servicesOnce = sync.Once{}
)

// GetServices Create a combined object for all services
func GetServices() Services {
	servicesOnce.Do(
		func() {
			// Init services here
			srvs.User = new(user.Service)
			srvs.Auth = new(auth.Service)
			srvs.Shared = shared.GetSharedServices()
		},
	)

	return *srvs
}
