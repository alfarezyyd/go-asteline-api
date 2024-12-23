package helper

import (
	"gorm.io/gorm"
)

func CheckErrorOperation(indicatedError error, errorForwarder func()) bool {
	if indicatedError != nil {
		errorForwarder()
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
