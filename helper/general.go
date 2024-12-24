package helper

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CheckErrorOperation(indicatedError error, ginContext *gin.Context, httpStatus int) bool {
	if indicatedError != nil {
		ginContext.AbortWithStatusJSON(httpStatus, gin.H{"error": indicatedError.Error()})
		return true
	}
	return false
}

func TransactionOperation(runningTransaction *gorm.DB, ginContext *gin.Context) {
	occurredError := recover()
	if occurredError != nil {
		runningTransaction.Rollback()
		ginContext.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": runningTransaction.Error.Error()})
	} else {
		runningTransaction.Commit()
	}
}
