package helper

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckErrorOperation(indicatedError error, ginContext *gin.Context, httpStatus int) bool {
	if indicatedError != nil {
		ginContext.AbortWithStatusJSON(httpStatus, gin.H{"error": indicatedError.Error()})
		return true
	}
	return false
}

func TransactionOperation(runningTransaction *gorm.DB) {
	occurredError := recover()
	if occurredError != nil {
		runningTransaction.Rollback()
	} else {
		runningTransaction.Commit()
	}
}
