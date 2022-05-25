package usecase

import "twitter_gin/internal/tweet/core/entity"

type tweetRespository interface {
	InsertTweet(tweet entity.Tweet, userID string) (string, error)
	GetTweet(ID string, page int64) ([]*entity.ResponseTweet, bool)
	DeleteTweet(ID string, userID string) error
}

type jwtService interface {
	GetUserID() string
}

type useCaseTweet struct {
	tweetReository tweetRespository
	jwtService     jwtService
}

func NewInsertUseCaseTweet(tweetReository tweetRespository, jwtService jwtService) *useCaseTweet {
	return &useCaseTweet{
		tweetReository: tweetReository,
		jwtService:     jwtService,
	}
}

func (us *useCaseTweet) InsertTweet(tweet entity.Tweet) (string, error) {
	userID := us.jwtService.GetUserID()
	return us.tweetReository.InsertTweet(tweet, userID)
}

func (us *useCaseTweet) GetTweet(ID string, page int64) ([]*entity.ResponseTweet, bool) {
	return us.tweetReository.GetTweet(ID, page)
}

func (us *useCaseTweet) DeleteTweet(ID string) error {
	userID := us.jwtService.GetUserID()
	return us.tweetReository.DeleteTweet(ID, userID)
}
