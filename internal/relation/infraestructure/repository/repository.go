package repository

import "twitter_gin/internal/relation/core/entity"

type repository interface {
	InsertRelation(relation entity.Relation) (bool, error)
	DeleteRelation(relation entity.Relation) (bool, error)
	GetRelation(relation entity.Relation) (bool, error)
	ListTweets(ID string, page int) ([]entity.ListTweets, bool)
}

type relationRepository struct {
	repository repository
}

func NewRelationRepository(repository repository) *relationRepository {
	return &relationRepository{
		repository: repository,
	}
}

func (ur relationRepository) InsertRelation(relation entity.Relation) (bool, error) {
	return ur.repository.InsertRelation(relation)
}
func (ur relationRepository) DeleteRelation(relation entity.Relation) (bool, error) {
	return ur.repository.DeleteRelation(relation)
}
func (ur relationRepository) GetRelation(relation entity.Relation) (bool, error) {
	return ur.repository.GetRelation(relation)
}
func (ur relationRepository) ListTweets(ID string, page int) ([]entity.ListTweets, bool) {
	return ur.repository.ListTweets(ID, page)
}
