package repository

import (
	"time"
	"twitter_gin/internal/tweet/core/entity"
)

type repository interface {
	InsertTweet(tweet entity.SaveTweet) (string, error)
	GetTweet(ID string, page int64) ([]*entity.ResponseTweet, bool)
	DeleteTweet(ID string, userID string) error
}

type tweetRepository struct {
	repository repository
}

func NewTweetRepository(repository repository) *tweetRepository {
	return &tweetRepository{
		repository: repository,
	}
}

func (tr tweetRepository) InsertTweet(tweet entity.Tweet, userID string) (string, error) {
	register := entity.SaveTweet{
		UserID:  userID,
		Message: tweet.Message,
		Date:    time.Now(),
	}
	return tr.repository.InsertTweet(register)
}

func (tr tweetRepository) GetTweet(ID string, page int64) ([]*entity.ResponseTweet, bool) {
	return tr.repository.GetTweet(ID, page)
}

func (tr tweetRepository) DeleteTweet(ID string, userID string) error {
	return tr.repository.DeleteTweet(ID, userID)
}
