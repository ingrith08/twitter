package user

import (
	"net/http"
	"strconv"
	"strings"
	"twitter_gin/internal/user/core/entity"

	"github.com/gin-gonic/gin"
)

type useCaseUser interface {
	Register(user entity.User) (string, error)
	Login(email, password string) (entity.User, bool)
	Profile(ID string) (entity.User, error)
	UpdateRegister(ID string, user entity.User) (bool, error)
	UploadAvatar(userID string, extension string) (bool, error)
	UploadBanner(userID string, extension string) (bool, error)
	ListUser(ID string, page int64, search string, typeUser string) ([]*entity.User, error)
}

type jwtService interface {
	CreateJWT(user entity.User) (string, error)
	GetUserID() string
}

type handler struct {
	usecase    useCaseUser
	jwtService jwtService
}

func NewInsetUserHandler(usecase useCaseUser, jwrjwtService jwtService) *handler {
	return &handler{
		usecase:    usecase,
		jwtService: jwrjwtService,
	}
}

func (h *handler) Register(ginCtx *gin.Context) {
	// Convertir al contexto de Meli
	var user entity.User

	if err := ginCtx.ShouldBind(&user); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(user.Email) == 0 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "El email de usuario es requerido"})
		return
	}

	if len(user.Password) < 6 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe especificar una contrasena de 6 caracteres"})
		return
	}

	id, err := h.usecase.Register(user)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, id)
}

func (h *handler) Login(ginCtx *gin.Context) {
	var user entity.User

	if err := ginCtx.ShouldBind(&user); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(user.Email) == 0 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "El email de usuario es requerido"})
		return
	}

	document, exist := h.usecase.Login(user.Email, user.Password)
	if !exist {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Usuario y/o contrasena invalidos"})
		return
	}

	jwtKey, err := h.jwtService.CreateJWT(document)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := entity.Login{
		Token: jwtKey,
	}

	ginCtx.JSON(http.StatusOK, token)
	ginCtx.SetCookie("token", jwtKey, 60*60*24, "/", "localhost", true, true)
}

func (h *handler) Profile(ginCtx *gin.Context) {
	ID := ginCtx.Param("id")
	if len(ID) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parametro ID"})
		return
	}
	user, err := h.usecase.Profile(ID)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, user)
}

func (h *handler) UpdateRegister(ginCtx *gin.Context) {
	ID := ginCtx.Param("id")
	if len(ID) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parametro ID"})
		return
	}

	var user entity.User
	if err := ginCtx.ShouldBind(&user); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := h.usecase.UpdateRegister(ID, user)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !status {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "No se ha logrado modificar el registro del usuario"})
		return
	}

	ginCtx.JSON(http.StatusOK, ID)
}

func (h *handler) UploadAvatar(ginCtx *gin.Context) {
	handler, err := ginCtx.FormFile("avatar")

	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := h.jwtService.GetUserID()
	var extension = strings.Split(handler.Filename, ".")[1]
	var dst string = "uploads/avatars/" + userID + "." + extension
	err = ginCtx.SaveUploadedFile(handler, dst)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := h.usecase.UploadAvatar(userID, extension)
	if err != nil || !status {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusCreated, "Succses")
}

func (h *handler) UploadBanner(ginCtx *gin.Context) {
	handler, err := ginCtx.FormFile("banner")
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := h.jwtService.GetUserID()
	var extension = strings.Split(handler.Filename, ".")[1]
	var dst string = "uploads/banners/" + userID + "." + extension
	err = ginCtx.SaveUploadedFile(handler, dst)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := h.usecase.UploadBanner(userID, extension)
	if err != nil || !status {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusCreated, "Succses")
}

func (h *handler) ListUser(ginCtx *gin.Context) {
	ID := h.jwtService.GetUserID()
	typeUser := ginCtx.Query("type")
	page := ginCtx.Query("page")
	search := ginCtx.Query("search")

	if len(page) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parámetro página"})
		return
	}
	if len(typeUser) < 1 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parámetro tipo"})
		return
	}
	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el parámetro página como entero mayor a 0"})
		return
	}
	pag := int64(pagTemp)
	users, err := h.usecase.ListUser(ID, pag, search, typeUser)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, users)
}
