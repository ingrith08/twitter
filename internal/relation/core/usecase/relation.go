package usecase

import "twitter_gin/internal/relation/core/entity"

type relationRespository interface {
	InsertRelation(relation entity.Relation) (bool, error)
	DeleteRelation(relation entity.Relation) (bool, error)
	GetRelation(relation entity.Relation) (bool, error)
	ListTweets(ID string, page int) ([]entity.ListTweets, bool)
}

type jwtService interface {
	GetUserID() string
}

type useCaseRelation struct {
	relationRespository relationRespository
	jwtService          jwtService
}

func NewInsertUseCaseRelation(relationRespository relationRespository, jwtService jwtService) *useCaseRelation {
	return &useCaseRelation{
		relationRespository: relationRespository,
		jwtService:          jwtService,
	}
}

func (us *useCaseRelation) InsertRelation(ID string) (bool, error) {
	var relation entity.Relation
	relation.UserID = us.jwtService.GetUserID()
	relation.UserRelationID = ID

	return us.relationRespository.InsertRelation(relation)
}
func (us *useCaseRelation) DeleteRelation(ID string) (bool, error) {
	var relation entity.Relation
	relation.UserID = us.jwtService.GetUserID()
	relation.UserRelationID = ID

	return us.relationRespository.DeleteRelation(relation)
}
func (us *useCaseRelation) GetRelation(ID string) (bool, error) {
	var relation entity.Relation
	relation.UserID = us.jwtService.GetUserID()
	relation.UserRelationID = ID

	return us.relationRespository.GetRelation(relation)
}
func (us *useCaseRelation) ListTweets(page int) ([]entity.ListTweets, bool) {
	ID := us.jwtService.GetUserID()
	return us.relationRespository.ListTweets(ID, page)
}
