package user

import "github.com/google/uuid"

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
	//
	return srv.repo.AddUser(user)
}

func (srv *UserService) DeleteUser(id uuid.UUID) error {
	return srv.repo.DeleteUser(id)
}
