package router

import "github.com/gin-gonic/gin"

func NewRouter(handlers *handlerContainer, router *gin.Engine) *gin.Engine {
	user := router.Group("/users")
	user.POST("/", handlers.userHandler.Register)
	user.POST("/login", handlers.userHandler.Login)
	user.GET("/:id", handlers.jwtMiddleware.ValidateJWT, handlers.userHandler.Profile)
	user.PUT("/:id", handlers.jwtMiddleware.ValidateJWT, handlers.userHandler.UpdateRegister)
	user.GET("/", handlers.jwtMiddleware.ValidateJWT, handlers.userHandler.ListUser)

	tweet := router.Group("/tweets")
	tweet.POST("/", handlers.jwtMiddleware.ValidateJWT, handlers.tweetHandler.SaveTweet)
	tweet.GET("/", handlers.jwtMiddleware.ValidateJWT, handlers.tweetHandler.GetTweet)
	tweet.DELETE("/:id", handlers.jwtMiddleware.ValidateJWT, handlers.tweetHandler.DeleteTweet)

	router.POST("/avatar", handlers.jwtMiddleware.ValidateJWT, handlers.userHandler.UploadAvatar)
	router.POST("/banner", handlers.jwtMiddleware.ValidateJWT, handlers.userHandler.UploadBanner)

	relation := router.Group("/relations")
	relation.POST("/", handlers.jwtMiddleware.ValidateJWT, handlers.relationHandler.CreateRelation)
	relation.DELETE("/:id", handlers.jwtMiddleware.ValidateJWT, handlers.relationHandler.DeleteRelation)
	relation.GET("/:id", handlers.jwtMiddleware.ValidateJWT, handlers.relationHandler.GetRelation)
	relation.GET("/", handlers.jwtMiddleware.ValidateJWT, handlers.relationHandler.ListTweets)
	return router
}
