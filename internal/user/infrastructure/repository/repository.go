package repository

import "twitter_gin/internal/user/core/entity"

type repository interface {
	InsertUser(user entity.User) (string, error)
	CheckUser(email string) (entity.User, bool)
	GetUser(ID string) (entity.User, error)
	UpdateUser(ID string, u map[string]interface{}) (bool, error)
	ListUser(page int64, search string) ([]*entity.User, error)
}

type userRepository struct {
	repository repository
}

func NewUserRepository(repository repository) *userRepository {
	return &userRepository{
		repository: repository,
	}
}

func (ur userRepository) InsertUser(user entity.User) (string, error) {
	return ur.repository.InsertUser(user)
}

func (ur userRepository) CheckUser(email string) (entity.User, bool) {
	return ur.repository.CheckUser(email)
}

func (ur userRepository) GetUser(ID string) (entity.User, error) {
	return ur.repository.GetUser(ID)
}

func (ur userRepository) UpdateUser(ID string, u entity.User) (bool, error) {
	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}
	register["date"] = u.Date
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		register["webSite"] = u.WebSite
	}
	return ur.repository.UpdateUser(ID, register)
}

func (ur userRepository) ListUser(page int64, search string) ([]*entity.User, error) {
	return ur.repository.ListUser(page, search)
}
