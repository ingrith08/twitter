package tweet

import (
	"net/http"
	"strconv"
	"twitter_gin/internal/tweet/core/entity"

	"github.com/gin-gonic/gin"
)

type useCaseTweet interface {
	InsertTweet(tweet entity.Tweet) (string, error)
	GetTweet(ID string, page int64) ([]*entity.ResponseTweet, bool)
	DeleteTweet(ID string) error
}

type handler struct {
	usecase useCaseTweet
}

func NewInsetTweetHandler(usecase useCaseTweet) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) SaveTweet(ginCtx *gin.Context) {
	var tweet entity.Tweet

	if err := ginCtx.ShouldBind(&tweet); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.usecase.InsertTweet(tweet)

	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusCreated, id)
}

func (h *handler) GetTweet(ginCtx *gin.Context) {
	ID := ginCtx.Query("id")
	page := ginCtx.Query("page")

	if len(ID) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parámetro id"})
		return
	}
	if len(page) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parámetro página"})
		return
	}
	pag, err := strconv.Atoi(page)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parámetro página con un valor mayor a 0"})
		return
	}

	pageInt := int64(pag)
	tweets, status := h.usecase.GetTweet(ID, pageInt)
	if !status {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer los tweets"})
		return
	}

	ginCtx.JSON(http.StatusOK, tweets)
}

func (h *handler) DeleteTweet(ginCtx *gin.Context) {
	ID := ginCtx.Param("id")
	if len(ID) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parametro ID"})
		return
	}

	err := h.usecase.DeleteTweet(ID)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, "Success delete")
}
