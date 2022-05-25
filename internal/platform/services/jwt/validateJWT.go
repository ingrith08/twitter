package jwt

import (
	"errors"
	"strings"
	"twitter_gin/internal/user/core/entity"

	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

func (service *jwtService) ValidateJWT(token string) (*entity.Claim, bool, string, error) {
	miClave := []byte("secret")
	claims := &entity.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, found := service.userReository.CheckUser(claims.Email)
		if found {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, found, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}
	return claims, false, string(""), err
}

func (service *jwtService) GetUserID() string {
	return IDUsuario
}
