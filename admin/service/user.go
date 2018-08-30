package service

import (
	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/repository"
)

// UserService is use action service
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService returns new instance
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// FetchUsers fetch users
func (u *UserService) FetchUsers(projectID string) ([]collection.User, error) {
	users, err := u.userRepo.Find(projectID)
	if err != nil {
		err = NewStoreSystemError(err.Error())
	}

	return users, err
}

// FetchUser fetch a user
func (u *UserService) FetchUser(uuid string, projectID string) (collection.User, error) {
	user, err := u.userRepo.FindByUUID(uuid, projectID)
	if err != nil {
		err = NewResourceNotFoundError(err.Error())
	}

	return user, err
}
