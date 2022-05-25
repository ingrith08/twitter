package router

import "github.com/gin-gonic/gin"

type handler interface {
	Register(ginCtx *gin.Context)
	Login(ginCtx *gin.Context)
	Profile(ginCtx *gin.Context)
	UpdateRegister(ginCtx *gin.Context)
	UploadAvatar(ginCtx *gin.Context)
	UploadBanner(ginCtx *gin.Context)
	ListUser(ginCtx *gin.Context)
}

type tweetHandler interface {
	SaveTweet(ginCtx *gin.Context)
	GetTweet(ginCtx *gin.Context)
	DeleteTweet(ginCtx *gin.Context)
}

type relationHandler interface {
	CreateRelation(ginCtx *gin.Context)
	DeleteRelation(ginCtx *gin.Context)
	GetRelation(ginCtx *gin.Context)
	ListTweets(ginCtx *gin.Context)
}

type middleware interface {
	ValidateJWT(ginCtx *gin.Context)
}

type handlerContainer struct {
	userHandler     handler
	tweetHandler    tweetHandler
	relationHandler relationHandler
	jwtMiddleware   middleware
}

func NewHandlerContainer(userHandler handler, tweetHandler tweetHandler, relationHandler relationHandler, jwtMiddleware middleware) *handlerContainer {
	return &handlerContainer{
		userHandler:     userHandler,
		tweetHandler:    tweetHandler,
		relationHandler: relationHandler,
		jwtMiddleware:   jwtMiddleware,
	}
}
