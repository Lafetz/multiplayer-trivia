package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUsernameUnique = errors.New("an account with this username exists")
	ErrDelete         = errors.New("failed to Delete user")
	ErrEmailUnique    = errors.New("an account with this email exists")
)

type UserService struct {
	repo repository
}

func NewUserService(repo repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (srv *UserService) GetUser(username string) (*User, error) {
	return srv.repo.GetUser(username)
}
func (srv *UserService) AddUser(user *User) (*User, error) {
	return srv.repo.AddUser(user)
}

func (srv *UserService) DeleteUser(id uuid.UUID) error {
	return srv.repo.DeleteUser(id)
}
