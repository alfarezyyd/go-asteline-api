package config

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func InitializeValidator() *validator.Validate {
	var validatorInstance = validator.New()
	// Tambahkan validasi kustom untuk 'datetime'
	validatorInstance.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		// Format tanggal yang valid
		format := "2006-01-02"
		_, err := time.Parse(format, fl.Field().String())
		return err == nil
	})
	return validatorInstance
}
