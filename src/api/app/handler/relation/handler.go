package relation

import (
	"net/http"
	"strconv"
	"twitter_gin/internal/relation/core/entity"

	"github.com/gin-gonic/gin"
)

type useCaseRelation interface {
	InsertRelation(ID string) (bool, error)
	DeleteRelation(ID string) (bool, error)
	GetRelation(ID string) (bool, error)
	ListTweets(page int) ([]entity.ListTweets, bool)
}

type handler struct {
	useCaseRelation useCaseRelation
}

func NewInsetRelationHandler(usecase useCaseRelation) *handler {
	return &handler{
		useCaseRelation: usecase,
	}
}

func (h *handler) CreateRelation(ginCtx *gin.Context) {
	var relation entity.RequestRelation

	if err := ginCtx.ShouldBind(&relation); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.useCaseRelation.InsertRelation(relation.UserRelationID)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusCreated, gin.H{})
}
func (h *handler) DeleteRelation(ginCtx *gin.Context) {
	ID := ginCtx.Param("id")
	if len(ID) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parametro ID"})
		return
	}

	_, err := h.useCaseRelation.DeleteRelation(ID)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, "Succses delete relation")
}

func (h *handler) GetRelation(ginCtx *gin.Context) {
	ID := ginCtx.Param("id")
	if len(ID) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parametro ID"})
		return
	}

	status, err := h.useCaseRelation.GetRelation(ID)
	if err != nil {
		status = false
	}

	ginCtx.JSON(http.StatusOK, gin.H{"status": status})
}

func (h *handler) ListTweets(ginCtx *gin.Context) {
	if len(ginCtx.Query("page")) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parámetro página"})
		return
	}
	page, err := strconv.Atoi(ginCtx.Query("page"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parámetro página con un valor mayor a 0"})
		return
	}

	response, status := h.useCaseRelation.ListTweets(page)
	if !status {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer los tweets"})
		return
	}

	ginCtx.JSON(http.StatusOK, response)
}
