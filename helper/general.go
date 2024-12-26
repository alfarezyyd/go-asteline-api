package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-asteline-api/exception"
	"gorm.io/gorm"
	"net/http"
)

func CheckErrorOperation(indicatedError error, ginContext *gin.Context, httpStatus int) bool {
	if indicatedError != nil {
		panic(exception.NewClientError(httpStatus, "Invalid request body"))
		return true
	}
	return false
}

func TransactionOperation(runningTransaction *gorm.DB, ginContext *gin.Context) {
	occurredError := recover()
	fmt.Println(occurredError)
	if occurredError != nil {
		fmt.Println(occurredError)
		runningTransaction.Rollback()
		ginContext.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": runningTransaction.Error.Error()})
	} else {
		runningTransaction.Commit()
	}
}
