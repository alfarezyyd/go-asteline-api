package category

import "github.com/gin-gonic/gin"

type Handler struct {
	categoryService *Service
}

func NewHandler(categoryService *Service) *Handler {
	return &Handler{
		categoryService: categoryService,
	}
}

func (categoryHandler *Handler) Create(ginContext *gin.Context) {
}

func (categoryHandler *Handler) Update(ginContext *gin.Context) {

}
func (categoryHandler *Handler) Delete(ginContext *gin.Context) {

}
