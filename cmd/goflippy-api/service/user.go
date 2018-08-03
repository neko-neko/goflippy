package service

import (
	"fmt"
	"time"

	"github.com/neko-neko/goflippy/collection"
	"github.com/neko-neko/goflippy/repository"
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

// RegisterUser register a user
func (u *UserService) RegisterUser(user *collection.User) error {
	_, err := u.userRepo.FindByUUID(user.UUID, user.ProjectID.Hex())
	if err == nil {
		return NewResourceAlreadyExistsError(fmt.Sprintf("user already exists %s", user.UUID))
	}
	if err := u.userRepo.Add(user); err != nil {
		return NewStoreSystemError(err.Error())
	}

	return nil
}

// UpdateUserInfo update a user info
func (u *UserService) UpdateUserInfo(user *collection.User) error {
	original, err := u.userRepo.FindByUUID(user.UUID, user.ProjectID.Hex())
	if err != nil {
		return NewResourceNotFoundError(fmt.Sprintf("user does not exists %s", user.UUID))
	}

	user.ID = original.ID
	user.CreatedAt = original.CreatedAt
	user.LastActivatedAt = original.LastActivatedAt
	user.UpdatedAt = time.Now()
	if err := u.userRepo.Update(user); err != nil {
		return NewStoreSystemError(err.Error())
	}

	return nil
}

// DeleteUser delete a user
func (u *UserService) DeleteUser(uuid string, projectID string) error {
	user, err := u.userRepo.FindByUUID(uuid, projectID)
	if err != nil {
		return NewResourceNotFoundError(fmt.Sprintf("user does not exists %s", user.UUID))
	}

	if err := u.userRepo.Delete(user.ID.Hex()); err != nil {
		return NewStoreSystemError(err.Error())
	}

	return nil
}
