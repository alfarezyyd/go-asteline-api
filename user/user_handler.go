package user

import "github.com/gin-gonic/gin"

type Handler struct {
	UserService *Service
}

func NewHandler(userService *Service) Controller {
	return &Handler{}
}

func (h *Handler) Register(ginContext *gin.Context) {

}
