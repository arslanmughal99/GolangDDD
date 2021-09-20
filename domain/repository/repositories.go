package repository

import (
	"sync"

	"massivleads/domain/repository/mongo"
	"massivleads/prototypes/repositories"
)

type Repositories struct {
	// Add Repositories type interfaces here
	UserRepository repositories.UserRepository
}

var (
	reposOnce = sync.Once{}
	repos     = new(Repositories)
)

func GetRepositories() Repositories {
	reposOnce.Do(
		func() {
			// Initialize all repositories here
			repos.UserRepository = mongo.NewMongoUserRepository()
		},
	)

	return *repos
}
