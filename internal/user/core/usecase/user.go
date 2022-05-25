package usecase

import (
	"errors"
	entityRelation "twitter_gin/internal/relation/core/entity"
	"twitter_gin/internal/user/core/entity"

	"golang.org/x/crypto/bcrypt"
)

type userRespository interface {
	InsertUser(user entity.User) (string, error)
	CheckUser(email string) (entity.User, bool)
	GetUser(ID string) (entity.User, error)
	UpdateUser(ID string, u entity.User) (bool, error)
	ListUser(ID string, page int64, search string, typeUser string) ([]*entity.User, error)
}

type relationRepository interface {
	GetRelation(relation entityRelation.Relation) (bool, error)
}

type useCaseUser struct {
	userReository      userRespository
	relationRepository relationRepository
}

func NewInsertUseCaseUser(userReository userRespository, relationRepository relationRepository) *useCaseUser {
	return &useCaseUser{
		userReository:      userReository,
		relationRepository: relationRepository,
	}
}

func (us *useCaseUser) Register(user entity.User) (string, error) {
	user.Password, _ = encriptarPassword(user.Password)
	_, found := us.userReository.CheckUser(user.Email)
	if found {
		return "", errors.New("ya existe usuario registrado")
	}
	return us.userReository.InsertUser(user)
}

func encriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}

func (us *useCaseUser) Login(email, password string) (entity.User, bool) {
	user, found := us.userReository.CheckUser(email)
	if !found {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}

func (us *useCaseUser) Profile(ID string) (entity.User, error) {
	return us.userReository.GetUser(ID)
}

func (us *useCaseUser) UpdateRegister(ID string, user entity.User) (bool, error) {
	return us.userReository.UpdateUser(ID, user)
}

func (us *useCaseUser) UploadAvatar(userID string, extension string) (bool, error) {
	var user entity.User
	user.Avatar = userID + "." + extension
	return us.userReository.UpdateUser(userID, user)
}

func (us *useCaseUser) UploadBanner(userID string, extension string) (bool, error) {
	var user entity.User
	user.Avatar = userID + "." + extension
	return us.userReository.UpdateUser(userID, user)
}

func (us *useCaseUser) ListUser(ID string, page int64, search string, typeUser string) ([]*entity.User, error) {
	users, err := us.userReository.ListUser(ID, page, search, typeUser)
	if err != nil {
		return users, err
	}

	var results []*entity.User

	for _, user := range users {
		var r entityRelation.Relation
		r.UserID = ID
		r.UserRelationID = user.ID.Hex()

		include := false
		found, _ := us.relationRepository.GetRelation(r)

		if typeUser == "new" && !found {
			include = true
		}
		if typeUser == "follow" && found {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}

		if include {

			results = append(results, user)
		}
	}

	return results, nil

}
