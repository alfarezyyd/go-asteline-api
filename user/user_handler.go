package user

import "github.com/gin-gonic/gin"

type Handler struct {
	UserService Service
}

func NewHandler(userService Service) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (h *Handler) Register(ginContext *gin.Context) {

}
