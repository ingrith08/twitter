package app

import (
	"log"
	"os"

	"twitter_gin/internal/platform/services/jwt"
	relationUseCase "twitter_gin/internal/relation/core/usecase"
	relationRepo "twitter_gin/internal/relation/infraestructure/repository"
	tweetUseCase "twitter_gin/internal/tweet/core/usecase"
	tweetRepo "twitter_gin/internal/tweet/infraestructure/repository"
	userUseCase "twitter_gin/internal/user/core/usecase"
	userRepo "twitter_gin/internal/user/infrastructure/repository"
	"twitter_gin/src/api/app/db"
	"twitter_gin/src/api/app/handler/relation"
	"twitter_gin/src/api/app/handler/tweet"
	"twitter_gin/src/api/app/handler/user"
	middleware "twitter_gin/src/api/app/middlewares"
	"twitter_gin/src/api/app/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartApp() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := gin.Default()
	URLDatabase := os.Getenv("URL_DATABASE")
	dbMongo := db.NewMongoRepository(URLDatabase, "twitter")

	userRepository := userRepo.NewUserRepository(dbMongo)
	tweetRepository := tweetRepo.NewTweetRepository(dbMongo)
	relationRepository := relationRepo.NewRelationRepository(dbMongo)

	jwtService := jwt.NewJwtService(userRepository)

	useCaseUser := userUseCase.NewInsertUseCaseUser(userRepository, relationRepository)
	useCaseTweet := tweetUseCase.NewInsertUseCaseTweet(tweetRepository, jwtService)
	useCaseRelation := relationUseCase.NewInsertUseCaseRelation(relationRepository, jwtService)

	userHandler := user.NewInsetUserHandler(useCaseUser, jwtService)
	tweetHandler := tweet.NewInsetTweetHandler(useCaseTweet)
	relationHandler := relation.NewInsetRelationHandler(useCaseRelation)
	jwtMiddleware := middleware.NewJwtMiddleware(jwtService)

	handlerContainer := router.NewHandlerContainer(userHandler, tweetHandler, relationHandler, jwtMiddleware)
	router := router.NewRouter(handlerContainer, server)

	return router.Run(os.Getenv("PORT"))
}
