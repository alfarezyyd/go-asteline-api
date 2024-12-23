package helper

import (
	"gorm.io/gorm"
)

func CheckErrorOperation(indicatedError error, errorForwarder func()) {
	if indicatedError != nil {
		errorForwarder()
	}
}

func TransactionOperation(runningTransaction *gorm.DB) {
	occurredError := recover()
	if occurredError != nil {
		runningTransaction.Rollback()
	} else {
		runningTransaction.Commit()
	}
}
