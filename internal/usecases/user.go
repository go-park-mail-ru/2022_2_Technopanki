package usecases

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/repository"
)

type UserService struct {
	ur repository.UserRepository
}

func newUserService(userRepos repository.UserRepository) *UserService {
	return &UserService{ur: userRepos}
}

func (us *UserService) CreateUser(user entity.User) error {
	return nil
}
func (us *UserService) GetUserByEmail(username string) (entity.User, error) {
	return entity.User{}, nil
}
