package middleware

import (
	"net/http"
	"twitter_gin/internal/user/core/entity"

	"github.com/gin-gonic/gin"
)

type jwtService interface {
	ValidateJWT(token string) (*entity.Claim, bool, string, error)
}

type jwtMiddleware struct {
	jwtService jwtService
}

func NewJwtMiddleware(jwtService jwtService) *jwtMiddleware {
	return &jwtMiddleware{
		jwtService: jwtService,
	}
}

func (m *jwtMiddleware) ValidateJWT(ginCtx *gin.Context) {
	_, _, _, err := m.jwtService.ValidateJWT(ginCtx.Request.Header["Authorization"][0])
	if err != nil {
		ginCtx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ginCtx.Next()
}
