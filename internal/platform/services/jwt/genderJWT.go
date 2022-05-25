package jwt

import (
	"time"
	"twitter_gin/internal/user/core/entity"

	jwt "github.com/dgrijalva/jwt-go"
)

type userRespository interface {
	CheckUser(email string) (entity.User, bool)
}

type jwtService struct {
	userReository userRespository
}

func NewJwtService(userReository userRespository) *jwtService {
	return &jwtService{
		userReository: userReository,
	}
}

func (service *jwtService) GenderJWT(user entity.User) (string, error) {
	myPassword := []byte("secret")

	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastName":  user.LastName,
		"date":      user.Date,
		"biography": user.Biography,
		"location":  user.Location,
		"webSite":   user.WebSite,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPassword)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
